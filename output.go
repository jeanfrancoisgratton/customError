// customError
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original timestamp: 2026/07/31 09:00
// Original filename: /output.go

package customError

import (
	"fmt"
	"io"
	"os"
)

// Error() only ever builds a string; deciding where that string goes is left to
// the caller. These helpers provide the sane default, so that every caller does
// not have to reinvent it (and, invariably, get it wrong by writing to stdout).

// Fprint renders the error to w.
//
// Colour is dropped automatically whenever w is not a terminal: redirected
// output, logfiles and shell command substitutions have no use for ANSI escape
// sequences. An explicit NoColourOutput = true is of course still honoured.
func (e CustomError) Fprint(w io.Writer) {
	if !isCharDevice(w) {
		e.NoColourOutput = true
	}
	fmt.Fprint(w, e.Error())
}

// Print renders the error to stderr.
//
// Diagnostics do not belong on stdout: stdout is the program's payload, and a
// caller doing  value=$(prog ...)  in a shell would otherwise capture the error
// message as though it were the value it asked for.
func (e CustomError) Print() {
	e.Fprint(os.Stderr)
}

// Die renders the error to stderr, then terminates the program with the exit
// code from ExitCode(). It is meant for the top-level caller (a cobra Run
// function, main(), ...) where there is nothing left to hand the error to.
//
// Note that Die() does not consult Fatality: asking to die is the caller's
// decision, whatever colour the message ends up being.
func (e CustomError) Die() {
	e.Print()
	os.Exit(e.ExitCode())
}

// ExitCode is the process exit status this error should translate to.
//
// PosixErrorCode wins when set, as that is precisely what it is there for.
// Failing that we fall back on Code, but only when it fits in the range a shell
// can report faithfully: POSIX exit statuses are 8-bit, and above 125 they
// collide with the values the shell reserves for its own signalling. Anything
// out of range, or no code at all, becomes a plain 1.
func (e CustomError) ExitCode() int {
	switch {
	case e.PosixErrorCode != 0:
		return e.PosixErrorCode
	case e.Code > 0 && e.Code < 126:
		return e.Code
	default:
		return 1
	}
}

// isCharDevice tells whether w is backed by a terminal. Anything that is not an
// *os.File (a bytes.Buffer in a test, a pipe wrapper, ...) is by definition not
// one, hence gets no colour.
func isCharDevice(w io.Writer) bool {
	f, ok := w.(*os.File)
	if !ok {
		return false
	}
	st, err := f.Stat()
	return err == nil && st.Mode()&os.ModeCharDevice != 0
}
