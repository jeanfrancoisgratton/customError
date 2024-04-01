package customError

// CustomError implements the error interface
// This will eventually be expanded

// The ErrortypeIota and the following constant are used to determine what kind of errors
// is CustomError able to manage
type ErrortypeIota int

const (
	Undefined ErrortypeIota = iota
	Fatal
	Warning
	Continuable
)

// This is the main data type.
type CustomError struct {
	Fatality ErrortypeIota // if omitted, we will use "Undefined, and will throw a panic() right away, so... define it :)
	Title    string        // optional
	Message  string        // if omitted, "Unspecified error" will be used
	Code     int           // optional
}

func (e CustomError) Error() string {
	if e.Message == "" {
		e.Message = "Unspecified errir"
	}
	switch e.Fatality {
	case ErrortypeIota(Fatal):
		return e.Fatal()
	case ErrortypeIota(Warning):
		return e.Warning()
	case ErrortypeIota(Continuable):
		return e.Continuable()
	default:
		e.Unknown()
	}
	return ""
}
