package main

import (
	"fmt"
	"math"
)

type Triplet [3]int

func checkNatural(a float64) bool{
	if a == float64(int64(a)) {
		return true
	}
	return false
}

func returnPythagoran(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

func main() {
	var sum int
	sum = 1000

	var TripletArray []Triplet
	
	for a := 1; a <= sum; a++ {
		for b := 1; a + b <= sum; b++ {
			c := returnPythagoran(a, b)
			if checkNatural(c) {
				// Check if sum property matches up
				if a + b + int(c) == sum && a < b && b < int(c) {
					TripletArray = append(TripletArray, Triplet{a, b, int(c)})
				}
			}
		}
	}
	
	fmt.Println(TripletArray)
}
