package main

import (
	"fmt"
	"unicode"
)

func RotationalCipher(plain string, shiftKey int) string {
	var result string
	var z rune
	for _, c := range plain {
		if  unicode.ToLower(c) >= 'a' &&  unicode.ToLower(c) <= 'z' {
			if unicode.IsLower(c) {
				z = 'z' 
			} else {
				z = 'Z'
			}
			// If letters, rotate
			if c + rune(shiftKey) > z {
				result += string(c + rune(shiftKey) - 26)
			} else {
				result += string(c + rune(shiftKey))
			}
		} else {
			// Otherwise, add in as is
			result += string(c)
		}
	}
	fmt.Println(result)
	return result
}

func main() {
	const inputShiftKey = 13
	var inputPlain string = "The quick brown fox jumps over the lazy dog."
	RotationalCipher(inputPlain, inputShiftKey)
}
