package tui

import "github.com/muesli/termenv"

// subtle returns a string with a subtle color.
func subtle(message string) string {
	return termenv.String(message).Foreground(termenv.ColorProfile().Color("241")).String()
}

// red returns a string with a red color.
func red(message string) string {
	return termenv.String(message).Foreground(termenv.ColorProfile().Color("196")).String()
}

// green returns a string with a green color.
func green(message string) string {
	return termenv.String(message).Foreground(termenv.ColorProfile().Color("46")).String()
}
