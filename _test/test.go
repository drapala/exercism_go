package main

import (
	"fmt"
	"regexp"
)

func main() {
	// Vowelsounds catcher
	vowel_word := "chair"
	results := regexp.MustCompile("a|e|i|o|u|xr|yt").FindAllStringIndex(vowel_word, -1)
	fmt.Println(results)

	// Use invert of vowelsounds catcher to get consonant clusters!

	// "qu" catcher - consonant only
	qu_word := "gkqueen"
	results = regexp.MustCompile("qu").FindAllStringIndex(qu_word, -1)
	fmt.Println(results)



	fmt.Println(regexp.MustCompile(" ").Split("gkqueen", -1))
	fmt.Println(regexp.MustCompile(" ").Split("quick fast run", -1))
}
