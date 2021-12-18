package main

import (
	"fmt"
	"unicode"
)

func CalculateCipher(char rune) rune {
	return rune(int('a') + int(rune('z') - char))
}

func main() {
	var s string = "Testing,1 2 3, testing."
	fmt.Println("Input: ", s)

	var output string
	var count int
	var lower rune

	for _, char := range(s) {
		lower = unicode.ToLower(char)
		// Only proceed for alphabets and numerics
		if (lower >= 'a' && lower <= 'z') || (lower >= '0' && lower <= '9') {
			// Group by 5
			if count == 5 {
				output += " "
				count = 0
			}
			if (lower >= 'a' && lower <= 'z') {
				// Deal with alphabets
				output += string(CalculateCipher(lower))
			} else if (lower >= '0' && lower <= '9') {
				// Deal with numerics
				output += string(char)
			}
			// Increment count
			count++
		}
	}
	fmt.Println("Output: ", output)
}
