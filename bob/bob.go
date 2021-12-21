package bob

import (
	"strings"
	"unicode"
)

func preProcess(s string) string {
	result := s
	// Remove leading and trailing spaces
	result = strings.TrimSpace(s)
	return result
}

func IsAllUpper(s string) bool {
	var isAllNumerics bool = true
	for _, r := range s {
		// If a single letter is not upper, is not all upper
		if !unicode.IsUpper(r) && unicode.IsLetter(r) {
			return false
		}
		// If we find a non-numeric character, is not all numerics
		if unicode.IsLetter(r) {
			isAllNumerics = false
		}
	}
	// If all numerics, return false since it's not uppercase
	if isAllNumerics {
		return false
	} else {
		return true
	}
}

func IsQuestion(s string) bool {
	if s[len(s)-1] == '?' {
		return true
	}
	return false
}

func Hey(remark string) string {
	// Pre-process the remark
	remark = preProcess(remark)

	// Check if silence
	if len(remark) == 0 {
		return "Fine. Be that way!"
	}
	// Check if question
	if IsQuestion(remark) {
		// Check if uppercase
		if IsAllUpper(remark) {
			return "Calm down, I know what I'm doing!"
		} else {
			return "Sure."
		}
	} else {
		// Check if uppercase
		if IsAllUpper(remark) {
			return "Whoa, chill out!"
		} else {
			return "Whatever."
		}
	}
}
