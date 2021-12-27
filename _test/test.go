package main

import (
	"fmt"
	"math"
)

func baseconv(n, base int) string {
	var result string
	for n > 0 {
		result = fmt.Sprintf("%d%s", n%base, result)
		n /= base
	}
	return result
}

func reverseArray(input []int) []int{
	output := make([]int, 0)
	for i:=len(input)-1; i>=0; i--{
		output = append(output, input[i])
	} 
	return output
}

func baseconvArray(n, base int) []int {
	result := make([]int, 0)
	for n > 0 {
		result = append(result, n%base)
		n /= base
	}
	return reverseArray(result)
}

func convertBaseTen(array []int, base int) int{
	var output int
	for i:=0; i<len(array); i++{
		// Basically, this is:
		// array[i] * base^position 
		output += array[i] * int(math.Pow(float64(base), math.Abs(float64(i-(len(array)-1)))))
	}
	return output
}

func main() {
	n:=42
	fmt.Println(baseconv(n, 2))
	fmt.Println(baseconvArray(n, 10))

	array_2 := []int{1,0,1,0,1,0}

	fmt.Println(array_2)
	fmt.Println(convertBaseTen(array_2, 2))
}