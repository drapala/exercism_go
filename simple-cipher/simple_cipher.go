package cipher

import (
	"math"
	"unicode"
)

// Define the shift and vigenere types here.
// Both types should satisfy the Cipher interface.
// Great tutorial on interfaces: https://gobyexample.com/interfaces

// Shift
type shift struct {
	distance int // Shift distance
}

// Normalizes rune into a-z range for both encode and decode
func shiftchar(newchar rune) string {
	var output string 
	if newchar > 'z' { // if we fall right-side of z
		output = string(newchar - 26) // wrap back left into a-z range
	} else if newchar < 'a' { // if we fall left-side of a
		output = string(newchar + 26) // wrap back right into a-z range
	} else {
		output = string(newchar) // keep as is
	}
	return output
}

// Cleans garbage out of input like spaces, non a-z etc.
func cleaninput(input string) string{
	var clean string
	for _, x := range input { // for each character in input
		x = unicode.ToLower(x) // normalize to lowercase
		if x >= 'a' && x <= 'z' { // only accepts a-z
			clean += string(x)
		}
	}
	return clean
}

// Note:
// 1. Encode will ignore all values in the string that are not A-Za-z.
// 2. The output will be also normalized to lowercase.
func (c shift) Encode(input string) string {
	var output string
	input = cleaninput(input)
	for _, x := range input { // for each character in clean input
		newchar := x + rune(c.distance) // Add since encode
		output += shiftchar(newchar)
	}
	return output
}

func (c shift) Decode(input string) string {
	var output string
	for _, x := range input { // for each character in input
		newchar := x - rune(c.distance) // Subtract since decode
		output += shiftchar(newchar)
	}
	return output
}

// Vigenere
type vigenere struct {
	key string
}

func propervignerekey(input, key string) string {
	var proper_key string = key
	var index int = 0
	for i:=1; i <= len(input) - len(key); i++ { // Loop for number of missing characters
		if index >= len(key) { // Wrap back
			index = 0
		}
		proper_key += string(key[index]) // Add repeated character to new key
		index++
	}
	return proper_key
}

func (v vigenere) Encode(input string) string {
	input = cleaninput(input) // Clean input so our key is clean

	// 1. Use the length of input to generate key with proper length
	proper_key := propervignerekey(input, v.key)

	// 2. Create a shift cipher per key, for per rune, and aggregate the output
	var output string
	for i, x := range(input) {
		if proper_key[i] == 'a' { // It's possible to have a zero shift as long as it's not ALL a's
			output += string(x)
		} else { // Non 'a' key
			c := NewShift(int(proper_key[i] - 'a')) // Set shift distance for cipher
			output += c.Encode(string(x)) // Add encoded character to output
		}
	}
	return output
}

func (v vigenere) Decode(input string) string {
	// 1. Use the length of input to generate key with proper length
	proper_key := propervignerekey(input, v.key)
	// 2. Create a shift cipher per key, for per rune, and aggregate the output
	var output string
	for i, x := range(input) {
		if proper_key[i] == 'a' { // It's possible to have a zero shift as long as it's not ALL a's
			output += string(x)
		} else { // Non 'a' key
			c := NewShift(int(proper_key[i] - 'a')) // Set shift distance for cipher
			output += c.Decode(string(x)) // Add encoded character to output
		}
	}
	return output
}


// Functions
func NewCaesar() Cipher {
	var c shift
	c.distance = 3
	return c
}

func NewShift(distance int) Cipher {
	var c shift
	if math.Abs(float64(distance)) <= 25 && math.Abs(float64(distance)) >= 1 { // Valid shift distance
		c.distance = distance
	} else { // Invalid shift distance
		return nil
	}
	return c
}

// Note:
// 1. Argument for NewVigenere must consist of lower case letters a-z only. 
// 2. Values consisting entirely of the letter 'a' are disallowed. 

// For invalid arguments NewVigenere returns nil.
func NewVigenere(key string) Cipher {
	var c vigenere

	var all_a bool = true // For keeping track of all a's 
	// 1. Argument for NewVigenere must consist of lower case letters a-z only. 
	for _, x := range(key) {
		if !unicode.IsLower(x) { // if not lowercase
			return nil
		}
		if x < 'a' || x > 'z' { // Not in a-z
			return nil
		}
		if x != 'a' { // Not all a's
			all_a = false
		}
	}
	// 2. Values consisting entirely of the letter 'a' are disallowed. 
	if all_a {
		return nil
	}
	// Key is valid
	c.key = key

	return c
}