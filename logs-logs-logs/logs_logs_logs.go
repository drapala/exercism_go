package logs

import (
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	recommendation, _ := utf8.DecodeRuneInString("❗")
	search, _ := utf8.DecodeRuneInString("🔍")
	weather, _ := utf8.DecodeRuneInString("☀")

	for _, rune := range log {
		switch rune {
			case recommendation:
				return "recommendation"
			case search:
				return "search"
			case weather:
				return "weather"
			default:
				continue;
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	// Covert string to array of runes to make mutable
	s := make([]rune, 0)

	// Append to array
	for _, rune := range log {
		if rune == oldRune {
			s = append(s, newRune)
		} else {
			s = append(s, rune)
		}
	}
	return string(s)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
