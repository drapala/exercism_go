package main

import "fmt"

func createZigZag(rails int, msgLength int) [][]string {
	rail_slice := make([][]string, 0)

	// initialize and append rails of appropriate lengths
	for i:=0; i < rails; i++{
		rail := make([]string, msgLength)
		rail_slice = append(rail_slice, rail)
	}

	var forward bool = true // Tracking direction
	var i int               // Tracking rail number

	for n := 0; n < msgLength; n++ {
		rail_slice[i][n] = "?"
		// Shift i
		if forward {
			i++
		} else {
			i--
		}
		// Change direction at end of rails
		if i == rails-1 {
			forward = false
		} else if i == 0 {
			forward = true
		}
	}
	// Print railslice
	for _, c := range(rail_slice) {
		fmt.Println(c)
	}

	return rail_slice
}

func main() {
	rails := 5
	msgLength := 17
	createZigZag(rails, msgLength)
}