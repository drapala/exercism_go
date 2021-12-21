package acronym

import (
	"strings"
)

func SplitString(s string) []string {
	var result []string
	// First, split on '-'
	for _, c := range(strings.Split(s, "-")){
		// Then, split on whitespace
		result = append(result, strings.Split(c, " ")...)
	}
	return result
}

func GetFirstAlphanumeric(s string) string {
	for _, c := range s {
		if c >= 'A' && c <= 'Z' {
			return string(c)
		}
	}
	return ""
}

func Abbreviate(s string) string {
	var result string
	// Split s on whitespace
	for _, word := range SplitString(s) {
		result += GetFirstAlphanumeric(strings.ToUpper(word))
	}
	return result
}
