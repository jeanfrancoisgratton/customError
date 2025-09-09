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
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s: %s\n", e.Red(e.Title), e.White(fmt.Sprintf("(error code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s: %s\n", e.Red(e.Title), e.White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s %s\n", e.Red(e.Message), e.White(fmt.Sprintf("error code: %v", strconv.Itoa(e.Code))))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s\n", e.Red(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Warning() string {
	builtString := ""
	if e.Title != "" {
		// with Title and Code
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s: %s\n", e.Yellow(e.Title), e.White(fmt.Sprintf("(error code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s: %s\n", e.Yellow(e.Title), e.White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s: %s\n", e.Yellow(e.Message), e.White(fmt.Sprintf("error code: %v", strconv.Itoa(e.Code))))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s\n", e.Yellow(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Continuable() string {
	builtString := ""
	if e.Title != "" {
		// with Title and Code
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s: %s\n", e.Blue(e.Title), e.White(fmt.Sprintf("(error code: %v) -> %s", strconv.Itoa(e.Code), e.Message)))
		} else {
			// With Title, no Code
			builtString = fmt.Sprintf("%s: %s\n", e.Blue(e.Title), e.White(e.Message))
		}
	} else {
		// No title, with code
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s: %s\n", e.Blue(e.Message), e.White(fmt.Sprintf("error code: %v", strconv.Itoa(e.Code))))
		} else {
			// No title, no code
			builtString = fmt.Sprintf("%s\n", e.Blue(e.Message))
		}
	}
	return builtString
}

// We do not care for Title, here...
func (e CustomError) Unknown() string {
	panic(fmt.Sprintf("\n\n%s\n", e.Message))
}
