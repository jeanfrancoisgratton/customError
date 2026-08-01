# Migrating callers to v3.1.0 reporting

## TL;DR

v3.1.0 is **strictly additive**. Bumping the dependency changes nothing at runtime and fixes nothing by itself. The old idiom still compiles, and neither the compiler nor `go vet` will flag it. Every call site has to be edited by hand.

## What went wrong, and why it is worth the sweep

Up to and including v3.0.0, the package only ever *built* a string; every caller picked its own sink, and the reflex was:

```go
fmt.Println(cerr.Error())
os.Exit(1)
```

That writes diagnostics to **stdout**. Stdout is the program's payload, so any caller doing

```bash
value=$(yourprogram --some-flag)
```

captures the error message *as the value it asked for* — and the result is not even empty, so the obvious defensive check (`[ -z "$value" ]`) passes and the garbage propagates.

This is not hypothetical. It took down Jenkins: `entrypoint.sh` read a keystore password out of Vault, Vault was sealed, and the skull-and-bones error text got handed to `--httpsKeyStorePassword`.

## The new API

```go
func (e CustomError) Fprint(w io.Writer)  // render to w
func (e CustomError) Print()              // render to stderr
func (e CustomError) Die()                // render to stderr, then os.Exit(ExitCode())
func (e CustomError) ExitCode() int       // PosixErrorCode, else Code (if 0 < Code < 126), else 1
```

`Fprint()` drops colour by itself whenever the destination is not a terminal, so redirected output and logfiles no longer need an explicit `NoColourOutput` (setting it yourself is still honoured).

## Per-tool checklist

1. `go get github.com/jeanfrancoisgratton/customError/v3@v3.1.0`
2. Grep the tool for the old idiom — do not assume, the compiler will not tell you:
   ```bash
   grep -rnE '(Print|Println|Printf)\(.*\.Error\(\)' --include='*.go' .
   grep -rn 'SkullBonesSign' --include='*.go' .
   ```
3. At the top level (`main()`, a cobra `Run`), collapse the whole error path to:
   ```go
   if cerr := doSomething(); cerr != nil {
       cerr.Die()
   }
   ```
4. Delete any intermediate printing. A helper that both prints *and* returns the error makes the message appear twice once the top level reports properly — return the error and let `Die()` be the single reporting point.
5. Populate `PosixErrorCode` (or a `Code` under 126) on error paths a script needs to tell apart, so callers see something better than a blanket `1`.
6. If the tool has a `--quiet` flag, make sure it suppresses the *payload* on stdout and never the diagnostics on stderr.
7. Verify from a shell, which is the only test that actually reproduces the original bug:
   ```bash
   out=$(yourprogram failing-args 2>/dev/null); echo "rc=$? captured=[$out]"
   # want: rc != 0, captured empty
   yourprogram failing-args 2>&1 >/dev/null | cat -v
   # want: the message, with no ANSI escapes
   ```

## Two traps

**The half-migration.** `fmt.Fprint(os.Stderr, cerr.Error())` gets the stream right but the colour wrong: gchalk probes *stdout* when deciding whether to emit escapes, so with stderr redirected into a logfile you still get ANSI codes whenever stdout happens to be a terminal. Only `Fprint`/`Print`/`Die` test the real destination. Use them rather than hand-rolling a sink.

**`Die()` ignores `Fatality`.** A `Warning` still exits the process if you call `Die()` on it. Asking to die is the caller's decision, whatever colour the message ends up being. Use `Print()` for anything the program should survive.

## Status

Inventory taken 2026-08-01. 24 repos depend on customError; the 17 below still print an error to stdout somewhere. The hit count is `grep -rEn '(Print|Println|Printf)\(.*\.Error\(\)'` and is an upper bound — some hits will be legitimate stdout output — but it is a fair measure of the work.

Note that the dependency version tells you nothing about whether a tool is migrated: `vclt` is already on v3.1.0 and is the single worst offender. That is the additive-release trap in one line.

| Done | Tool | customError | hits |
|------|------|-------------|------|
| [x] | devops/vaultreader | v3.1.0 | 0 |
| [ ] | devops/vclt | v3.1.0 | 46 |
| [ ] | devops/pgtools | **v2.3.3** | 43 |
| [ ] | devops/nxtools | v3.0.0 | 22 |
| [ ] | devops/dvol | v3.0.0 | 15 |
| [ ] | devops/vmman4 | v3.0.0 | 14 |
| [ ] | devops/monctl | v3.0.0 | 12 |
| [ ] | devops/certificatemanager | v3.0.0 | 8 |
| [ ] | devops/dtools2 | v3.0.0 | 8 |
| [ ] | devops/prometheusConsumer | v3.0.0 | 5 |
| [ ] | mainline/encdec | v3.0.0 | 4 |
| [ ] | devops/prometheusListener | v3.0.0 | 2 |
| [ ] | devops/purgePackages | (indirect) | 2 |
| [ ] | mainline/stubber | v3.0.0 | 1 |
| — | vl_oldtimers/certificatemanager | (indirect) | 8 |
| — | vl_oldtimers/dvol | v3.0.0 | 15 |
| — | vl_oldtimers/dtools2 | v3.0.0 | 8 |
| — | vl_oldtimers/stubber | v3.0.0 | 1 |

`vl_oldtimers/*` are marked `—` on the assumption that they are retired; promote them if not.

Clean already, no stdout printing found: `aylo/assetmgr`, `devops/buildlistener`, `mainline/batchFileDelete`, `mainline/dircmp`, `packages/giteaops`, `packages/hcpVaultLib`.

`devops/pgtools` is still on v2.3.3, so it needs the `/v3` module path bump first — see the 3.0.0 CHANGELOG row.

### Prior art

`devops/vaultreader` is migrated and is the reference implementation: `cmd/root.go` collapses to `kvreadErr.Die()`, the intermediate `fmt.Println(hftx.SkullBonesSign(...))` blocks in `kv/helpers.go` are gone, and the sealed-Vault path sets a real error code so the calling script can distinguish "retry, it is only sealed" from "bad token, give up". Copy that shape.
