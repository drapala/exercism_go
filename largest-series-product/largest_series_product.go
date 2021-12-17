package lsproduct

import (
	"fmt"
	"strconv"
)

func CheckValidNumber(digit string) bool {
	for i := 0; i < len(digit); i++ {
		if digit[i] < '0' || digit[i] > '9' {
			return false
		}
	}
	return true
} 

func CalculateProduct(subdigit string) int {
	var result int = 1
	for _, v := range subdigit {
		conv, _ := strconv.Atoi(string(v))
		result = result * conv
	}
	return result
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	// Error handling
	if span > len(digits) {
		return 0, fmt.Errorf("span must be smaller than string length")
	} else if span < 0 {
		return 0, fmt.Errorf("span must not be negative")
	} else if !CheckValidNumber(digits) {
		return 0, fmt.Errorf("digits input must only contain digits")
	}
	// Number of possible permutations
	perm := len(digits) - span + 1
	var max_product int = 0
	for i := 0; i < perm; i++ {
		product := CalculateProduct(digits[i:i+span])
		if max_product < product {
			max_product = product
		}
	}
	return int64(max_product), nil
}
