package main

import (
	"fmt"
	"strconv"
)

func CalculateProduct(subdigit string) int {
	var result int = 1
	for _, v := range subdigit {
		conv, _ := strconv.Atoi(string(v))
		result = result * conv
	}
	return result
}

func main() {
	digits := "1027839564"
	span := 5
	perm := len(digits) - span + 1

	var max_product int = 0
	for i := 0; i < perm; i++ {
		product := CalculateProduct(digits[i:i+span])
		if max_product < product {
			max_product = product
		}
	}
	fmt.Println("Max product is: ", max_product)
}
