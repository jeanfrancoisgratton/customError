# customError

A small Go package extending the `error` interface with coloured, titled, coded error output.

The source of truth is the gitea repo (`origin`); `gh` is a push-only mirror to github. Never `git pull` from `gh`.

## Reporting errors — read before touching any consumer

`Error()` only *builds* a string. Since v3.1.0 the package owns the sink as well:
`Fprint(w)`, `Print()` (stderr), `Die()` (stderr + `os.Exit(ExitCode())`), `ExitCode()`.

**Diagnostics go to stderr, never stdout.** A tool that does `fmt.Println(cerr.Error())` will have its error text captured as a value by any shell doing `value=$(tool ...)` — non-empty garbage that defeats the obvious `[ -z "$value" ]` check. This took down a Jenkins container in 2026-08.

v3.1.0 is strictly additive, so **bumping the dependency fixes nothing** and the compiler will never flag the old idiom. A migration sweep across ~17 consumer repos is in progress — see [docs/MIGRATION-v3.1.0.md](docs/MIGRATION-v3.1.0.md) for the checklist, the two traps, and the per-tool status table. `devops/vaultreader` is the reference implementation.

## Conventions

- `current_package_version` and `go.version` are plain-text version files; keep them in step with `docs/CHANGELOG.md`, which is a markdown table, newest release on top.
- Every release gets a CHANGELOG row. 3.0.0 was missed and had to be reconstructed from the tagged commit; do not repeat that.
- Colour helpers are methods on `CustomError` so they can honour `NoColourOutput`; keep new output paths going through them rather than calling gchalk directly.
