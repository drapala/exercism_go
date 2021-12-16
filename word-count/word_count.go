package wordcount

import (
	"strings"
)

type Frequency map[string]int

func strip(s string) string {
    var result strings.Builder
    for i := 0; i < len(s); i++ {
        b := s[i]
        if ('a' <= b && b <= 'z') ||
            ('A' <= b && b <= 'Z') ||
            ('0' <= b && b <= '9') ||
            b == ' ' ||
			b == '\''{
            result.WriteByte(b)
        }
    }
    return result.String()
}

func CleanWord(in string) string {
	var out string

	// Skip empty strings
	if in == "" {
		return ""
	}
	// Skip punctuation
	out = strip(in)

	// Deal with quotations in beginning and end
	if out[0] == '\'' && out[len(out)-1] == '\'' {
		out = out[1:len(out)-1]
	}

	return out
}

func WordCount(phrase string) Frequency {
	result := make(Frequency)

	// Split by whitespace first
	// https://pkg.go.dev/strings#Fields
	whitesplit := strings.Fields(phrase)

	// Split per entry by specific delimiters
	delimited := make([]string, 0)
	// https://pkg.go.dev/strings#Split
	for _, entry := range whitesplit {
		for _, word := range strings.Split(entry, ",") {
			// Append to delimited
			// Filter stuff that didn't get caught:
			if CleanWord(word) != "" {
				delimited = append(delimited, CleanWord(word))
			}
		}
	}

	for _, word := range(delimited) {
		result[strings.ToLower(word)]++
	}

	return result
}
