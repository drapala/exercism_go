package main

import (
	"fmt"
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
    for _, r := range s {
        if !unicode.IsUpper(r) && unicode.IsLetter(r) {
            return false
        }
    }
    return true
}

func IsQuestion(s string) bool {
	if s[len(s)-1] == '?' {
		return true
	}
    return false
}

func main() {
	s1 := "   Tom-ay-to, tom-aaaah-to.  "
	s1 = preProcess(s1)
	fmt.Println("s1 is Uppercase?", IsAllUpper(s1))

	s2 := "  WHAT'S GOING ON?   "
	s2 = preProcess(s2)
	fmt.Println("s2 is question?", IsQuestion(s2))
	fmt.Println("s2 is Uppercase?", IsAllUpper(s2))
}
