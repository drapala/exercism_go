package main

import "fmt"

type nums struct {
	value int
	numtype string
}

func generateNums(limit int) []nums {
	numArrays := make([]nums, 0)
	for i := 2; i <= limit ; i++ {
		currentNum := nums{value: i, numtype: "tbd"}
		numArrays = append(numArrays, currentNum)
	}
	return numArrays
}

func filterprimes(numArray []nums) []int {
	primes := make([]int, 0)
	for i:=0; i < len(numArray); i++ {
		if numArray[i].numtype == "prime" {
			primes = append(primes, numArray[i].value)
		}
	}
	return primes
}

func main() {
	var limit int = 13

	var numArray []nums
	numArray = generateNums(limit)

	for i := 0; i < len(numArray); i++ {
		currentVal := numArray[i].value
		if numArray[i].numtype == "composite"{
			continue
		} else if numArray[i].numtype == "tbd" {
			// Set current to prime
			numArray[i].numtype = "prime"
			// index = currentVal - 2
			// So we add currentVal to get the next multiple's index
			nextIndex := (currentVal - 2) + currentVal
			// While loop to mark the rest of the multiples as composites
			for nextIndex < len(numArray) {
				numArray[nextIndex].numtype = "composite"
				nextIndex += currentVal
			}
		}			
	}
	fmt.Println(numArray)

	primes := filterprimes(numArray)
	fmt.Println(primes)

}
