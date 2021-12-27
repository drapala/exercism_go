package allyourbase

import (
	"fmt"
	"math"
)

// Reverses an array
func reverseArray(input []int) []int{
	output := make([]int, 0)
	for i:=len(input)-1; i>=0; i--{
		output = append(output, input[i])
	} 
	return output
}
// Converts an integer in base10 to a given base
func baseconvArray(n, base int) []int {
	result := make([]int, 0)
	for n > 0 {
		result = append(result, n%base)
		n /= base
	}
	if n == 0 && len(result) == 0 { // We need to return a single 0 array
		return []int{0}
	}
	return reverseArray(result)
}

// Converts a number from a given base to base10
func convertBaseTen(array []int, base int) int{
	var output int
	for i:=0; i<len(array); i++{
		// Basically, this is:
		// array[i] * base^position 
		output += array[i] * int(math.Pow(float64(base), math.Abs(float64(i-(len(array)-1)))))
	}
	return output
}

// Returns true if 0 <= d < input base is satisfied
func baseValidChecker(array []int, base int) bool {
	var valid bool = true
	for i:=0; i<len(array); i++{
		if array[i] < 0 || array[i] >= base {
			valid = false
		}
	}
	return valid
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	// Error handling
	if inputBase < 2 {
		return nil, fmt.Errorf("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, fmt.Errorf("output base must be >= 2")
	}
	if !baseValidChecker(inputDigits, inputBase){
		return nil, fmt.Errorf("all digits must satisfy 0 <= d < input base")
	}
	// Edge case handling
	if len(inputDigits) == 0 { // Empty list
		return []int{0}, nil
	}
	num_base_10 := convertBaseTen(inputDigits, inputBase) // Normalize to base10 first
	outputDigits := baseconvArray(num_base_10, outputBase) // Convert from base10 to desired base
	return outputDigits, nil
}
