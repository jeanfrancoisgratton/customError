// customErrors
// Written by J.F. Gratton <jean-francois@famillegratton.net>
// Original filename: customErrors/terminalFx.go
// Original timestamp: 2024/03/30 14:51

package customError

import (
	"fmt"
	"github.com/jwalton/gchalk"
)

// COLOR FUNCTIONS
// ===============
func Red(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightRed().Bold(sentence))
}

func Green(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightGreen().Bold(sentence))
}

func White(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightWhite().Bold(sentence))
}

func Yellow(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightYellow().Bold(sentence))
}

func Blue(sentence string) string {
	return fmt.Sprintf("%s", gchalk.WithBrightBlue().Bold(sentence))
}
