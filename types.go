// customError
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original timestamp: 2025/09/26 14:22
// Original filename: /types.go

package customError

// This is the main data type.
type CustomError struct {
	Fatality       ErrortypeIota // optional, if omitted, we will use "Fatal"; "Undefined" will throw a panic right away
	Title          string        // optional
	Message        string        // optional, if omitted, "Unspecified error" will be used
	Code           int           // optional
	PosixErrorCode int           // Error codes above 128 are not well handled in POSIX, we will allow mapping for POSIX codes to be returned with os.Exit()
	NoColourOutput bool          // handles if we use colours in output (useless in logfiles)
}
