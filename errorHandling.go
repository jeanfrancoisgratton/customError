package customError

import (
	"fmt"
	"strconv"
)

// This is how a FATAL error message is displayed, with various scenarios:
// - with or without a Title
// - with or without an error code
func (e CustomError) Fatal() string {
	builtString := ""

	if e.Title != "" {
		// with Title and Code
		e.Title += ":"
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s %s\n", Red(e.Title), White(fmt.Sprintf("(code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s %s\n", Red(e.Title), White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s %s\n", Red(e.Title), White(fmt.Sprintf("(code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s: %s\n", Red(e.Title), White(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Warning() string {
	builtString := ""
	if e.Title != "" {
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s: %s\n", Yellow(e.Title), White(fmt.Sprintf("(code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Yellow(e.Title), White(e.Message))
		}
	} else {
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s (code: %d)\n", Yellow(e.Message), White(fmt.Sprintf("(code: %d)\n", strconv.Itoa(e.Code))))
		} else {
			builtString = fmt.Sprintf("%s\n", Yellow(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Continuable() string {
	builtString := ""
	if e.Title != "" {
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s (code: %d): %s\n", Blue(e.Title), White(strconv.Itoa(e.Code)), Blue(e.Message))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Blue(e.Title), White(e.Message))
		}
	} else {
		if e.Code != 0 {

			builtString = fmt.Sprintf("%s (code: %d): %s\n", Blue(e.Message), White(strconv.Itoa(e.Code)))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Blue(e.Title), White(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Unknown() string {
	panic(fmt.Sprintf("\n\n%s\n", e.Message))
}
