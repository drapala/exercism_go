package main

import "fmt"

func getDivisors(n int) []int {
	result := make([]int, 0)
	for x := 1; x < n; x++ { // Find divisors
		if n % x == 0 { // Divides cleanly
			result = append(result, x)
		}
	}
	return result
}

func getDivisorSum(array []int) int {
	var sum int
	for _, divisor := range(array) {
		sum += divisor
	}
	return sum
}

func main() {
	n := 28
	fmt.Println(n)
	fmt.Println(getDivisors(n))
	fmt.Println(getDivisorSum(getDivisors(n)))
}