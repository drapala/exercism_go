package main

import "fmt"

func main() {
	num1 := 1234567890

	billions := num1 / 1000000000
	millions := num1 / 1000000
	thousands := num1 / 1000

	fmt.Println(billions) // Billion = 1
	fmt.Println(millions - 1000 * billions) // Million = 234
	fmt.Println(thousands - 1000 * millions) // Thousand = 567
	fmt.Println(num1 - 1000 * thousands) // Hundred = 890

	fmt.Println("==================")

	num2 := 234567890

	billions = num2 / 1000000000
	millions = num2 / 1000000
	thousands = num2 / 1000

	fmt.Println(billions) // Billion = 0
	fmt.Println(millions - 1000 * billions) // Million = 234
	fmt.Println(thousands - 1000 * millions) // Thousand = 567
	fmt.Println(num2 - 1000 * thousands) // Hundred = 890
}
