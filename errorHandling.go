package customError

import (
	"fmt"
	"os"
	"strconv"
)

// This is how a FATAL error message is displayed, with various scenarios:
// - with or without a Title
// - with or without an error code
func (e CustomError) Fatal() {
	builtString := ""
	if e.Title != "" {
		// with Title and Code
		e.Title += ":"
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s %s\n", Red(e.Title), White(fmt.Sprintf("(error code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s %s\n", Red(e.Title), White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			e.Message += ":"
			builtString = fmt.Sprintf("%s %s\n", Red(e.Message), White(fmt.Sprintf("error code: %v", strconv.Itoa(e.Code))))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s\n", Red(e.Message))
		}
	}
	fmt.Println(builtString)
	os.Exit(e.Code)
}

func (e CustomError) Warning() string {
	builtString := ""
	if e.Title != "" {
		// with Title and Code
		e.Title += ":"
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s %s\n", Yellow(e.Title), White(fmt.Sprintf("(error code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s %s\n", Yellow(e.Title), White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			e.Message += ":"
			builtString = fmt.Sprintf("%s %s\n", Yellow(e.Message), White(fmt.Sprintf("error code: %v", strconv.Itoa(e.Code))))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s\n", Yellow(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Continuable() string {
	builtString := ""
	if e.Title != "" {
		// with Title and Code
		e.Title += ":"
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s %s\n", Blue(e.Title), White(fmt.Sprintf("(error code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s %s\n", Blue(e.Title), White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			e.Message += ":"
			builtString = fmt.Sprintf("%s %s\n", Blue(e.Message), White(fmt.Sprintf("error code: %v", strconv.Itoa(e.Code))))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s\n", Blue(e.Message))
		}
	}
	return builtString
}

// We do not care for Title, here...
func (e CustomError) Unknown() string {
	panic(fmt.Sprintf("\n\n%s\n", e.Message))
}
