package customError

import (
	"fmt"
	"strconv"
)

func (e CustomError) Fatal() string {
	builtString := ""

	if e.Title != "" {
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s (code: %d): %s\n", Red(e.Title), White(strconv.Itoa(e.Code)), Red(e.Message))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Red(e.Title), White(e.Message))
		}
	} else {
		if e.Code != 0 {

			builtString = fmt.Sprintf("%s (code: %d): %s\n", Red(e.Message), White(strconv.Itoa(e.
				Code)))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Red(e.Title), White(e.Message))
		}
	}
	return builtString
}

func (e CustomError) Warning() string {
	builtString := ""
	if e.Title != "" {
		if e.Code != 0 {
			builtString = fmt.Sprintf("%s (code: %d): %s\n", Yellow(e.Title), White(strconv.Itoa(e.Code)), Yellow(e.Message))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Yellow(e.Title), White(e.Message))
		}
	} else {
		if e.Code != 0 {

			builtString = fmt.Sprintf("%s (code: %d): %s\n", Yellow(e.Message), White(strconv.Itoa(e.Code)))
		} else {
			builtString = fmt.Sprintf("%s: %s\n", Yellow(e.Title), White(e.Message))
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
