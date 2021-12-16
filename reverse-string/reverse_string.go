package reverse

import "unicode/utf8"

// https://go.dev/blog/strings
func Reverse(input string) string {
	// Store runes in a slice
	store := []rune{}
	for i, w := 0, 0; i < len(input); i += w {
		runeValue, width := utf8.DecodeRuneInString(input[i:])
		w = width
		store = append(store, runeValue)
	}
	// Reverse and return result
	var result string
	for i := len(store) - 1; i >= 0; i-- {
		result += string(store[i])
	}
	return result
}
