package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	const input = "Hello, 世界"

	store := []rune{}
	for i, w := 0, 0; i < len(input); i += w {
		runeValue, width := utf8.DecodeRuneInString(input[i:])
		w = width
		store = append(store, runeValue)
	}

	var result string

	for i := len(store) - 1; i >= 0; i-- {
		result += string(store[i])
	}
	fmt.Println(result)
}
