package pythagorean

import (
	"math"
)

type Triplet [3]int

// Checks if a number is natural
func checkNatural(a float64) bool {
	if a == float64(int64(a)) {
		return true
	}
	return false
}

// Given 2 sides, returns the third side of the pythagorean triplet
func returnPythagoran(a, b int) float64 {
	return math.Sqrt(float64(a*a + b*b))
}

// Range returns a list of all Pythagorean triplets with sides in the
// range min to max inclusive.
func Range(min, max int) []Triplet {
	var TripletArray []Triplet
	for a := min; a <= max; a++ {
		for b := a + 1; b <= max; b++ {
			c := returnPythagoran(a, b)
			if checkNatural(c) && int(c) <= max {
				TripletArray = append(TripletArray, Triplet{a, b, int(c)})
			}
		}
	}
	return TripletArray
}

// Sum returns a list of all Pythagorean triplets where the sum a+b+c
// (the perimeter) is equal to p.
// The three elements of each returned triplet must be in order,
// t[0] <= t[1] <= t[2], and the list of triplets must be in lexicographic
// order.
func Sum(p int) []Triplet {
	var TripletArray []Triplet
	for a := 1; a <= p; a++ {
		for b := 1; a + b <= p; b++ {
			c := returnPythagoran(a, b)
			if checkNatural(c) {
				// Check if sum property matches up
				if a + b + int(c) == p && a < b && b < int(c) {
					TripletArray = append(TripletArray, Triplet{a, b, int(c)})
				}
			}
		}
	}
	return TripletArray
}
