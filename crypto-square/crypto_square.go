package cryptosquare

import (
	"unicode"
)

func normalize(in string) string {
	var out string
	for _, c := range in {
		if unicode.IsLetter(c) || unicode.IsNumber(c) {
			out += string(unicode.ToLower(c))
		}
	}
	return out
}

func findRectangleSpecs(area int) (int, int) {
	for c := 1; c <= area; c++ {
		for r := 1; r <= c; r++ {
			if (r*c >= area) && (c-r <= 1) {
				return r, c
			}
		}
	}
	// Couldn't find anything for some reason
	return -1, -1
}

func padWhitespace(text string, c int) string {
	var out string = text
	for i := len(text); i < c; i++ {
		out += " "
	}
	return out
}

func createRectangle(text string, c int) []string {
	rectangle := make([]string, 0)
	var counter int
	var temp string
	for i := 0; i < len(text); i++ {
		temp += string(text[i])
		counter++
		// If width is filled up, append block as row and reset for new row
		if counter == c {
			rectangle = append(rectangle, temp)
			temp = ""
			counter = 0
		}
	}
	// Pad and append left over
	if temp != "" {
		rectangle = append(rectangle, padWhitespace(temp, c))
	}
	return rectangle
}

func createCipherFromRectangle(rectangle []string) string {
	var cipher string
	r := len(rectangle)
	c := len(rectangle[0])
	for col := 0; col < c; col++ {
		for row := 0; row < r; row++ {
			cipher += string(rectangle[row][col])
			if (row == r-1) && (col != c-1) {
				// If we're not at the last entry, append whitespace
				cipher += " "
			}
		}
	}
	return cipher
}

func Encode(pt string) string {
	// Edge case
	if len(pt) == 0 {
		return pt
	}
	// Normalize string
	var normalized string = normalize(pt)
	// Determine rectangle specs
	var c int
	_, c =  findRectangleSpecs(len(normalized))
	// Fill rectangle
	var rectangle []string
	rectangle = createRectangle(normalized, c)
	// Generate cipher from rectange
	var cipher string
	cipher = createCipherFromRectangle(rectangle)
	return cipher
}
