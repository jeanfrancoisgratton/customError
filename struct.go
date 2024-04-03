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
	Fatality       ErrortypeIota // optional, if omitted, we will use "Undefined, and will throw a panic() right away, so... define it :)
	Title          string        // optional
	Message        string        // optional, if omitted, "Unspecified error" will be used
	Code           int           // optional
	ColouredOutput bool          // unused for now
}

// This is the main function, really
func (e CustomError) Error() string {
	if e.Message == "" {
		e.Message = "Unspecified error"
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

// THIS IS A CONSTRUCTOR STUB !!!
// UNUSED FOR NOW
// CustomError constructor :
// For now, we only need to ensure that the "ColouredOuput" struct member is initialized, that is, passed in parameters

func NewCustomError(colouredOutput bool) CustomError {
	// for now, not doing anything

	return CustomError{}
}
