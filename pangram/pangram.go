package pangram

import (
	"strings"
)

func IsPangram(input string) bool {

	// turn input to lowercase
	input = strings.ToLower(input)

	// To store flags
	flagger := make(map[rune]bool, 123-97)

	// Loop over string and store flags
	for _, r := range input {
		flagger[r] = true
	}

	var r rune
	for i := 97; i <= 122; i++ {
		// Convert i to corresponding rune
		r = rune(i)
		if flagger[r] == false {
			return false
		}
	}
	return true
}
