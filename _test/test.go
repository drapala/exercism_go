package main

import (
	"fmt"
	"regexp"
)

func main() {
	input := "What is -3 plus 7 multiplied by -2?"
    inter := regexp.MustCompile(`What is `).ReplaceAllString(input, "")
	output := regexp.MustCompile(`\?`).ReplaceAllString(inter, "")

	fmt.Println(output)
}
