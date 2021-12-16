package main

import "fmt"

func main() {
	const input = "'large'"
	fmt.Println(input)
	fmt.Println(input[1:len(input)-1])
}
