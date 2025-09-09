// customError
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: customError/terminalFx.go
// Original timestamp: 2024/03/30 14:51

package customError

import (
	"fmt"
	"github.com/jwalton/gchalk"
)

// COLOR FUNCTIONS
// ===============
func (e CustomError) Red(sentence string) string {
	if !e.NoColourOutput {
		return fmt.Sprintf("%s", gchalk.WithBrightRed().Bold(sentence))
	}
	return sentence
}

func (e CustomError) Green(sentence string) string {
	if !e.NoColourOutput {
		return fmt.Sprintf("%s", gchalk.WithBrightGreen().Bold(sentence))
	}
	return sentence
}

func (e CustomError) White(sentence string) string {
	if !e.NoColourOutput {
		return fmt.Sprintf("%s", gchalk.WithBrightWhite().Bold(sentence))
	}
	return sentence
}

func (e CustomError) Yellow(sentence string) string {
	if !e.NoColourOutput {
		return fmt.Sprintf("%s", gchalk.WithBrightYellow().Bold(sentence))
	}
	return sentence
}

func (e CustomError) Blue(sentence string) string {
	if !e.NoColourOutput {
		return fmt.Sprintf("%s", gchalk.WithBrightBlue().Bold(sentence))
	}
	return sentence
}

func (e CustomError) ClearTerminal() {
	fmt.Print("\033[2J\033[H") // ANSI escape code for clearing the terminal
}
