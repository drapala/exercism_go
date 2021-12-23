package main

import "fmt"

func contains(s []rune, e rune) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	// input := "\\left(\\begin{array}{cc} \\frac{1}{3} & x\\\\ \\mathrm{e}^{x} &... x^2 \\end{array}\\right)"
	input := "(((185 + 223.85) * 15) - 543)/2"

	var cleaned string
	match := []rune{'[', ']', '(', ')', '{', '}'}

	for _, c := range input {
		if contains(match, c) {
			cleaned += string(c)
		}
	}
	fmt.Println(cleaned)
}
