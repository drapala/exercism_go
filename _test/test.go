package main

import (
	"fmt"
	"math"
)

func getMax1Bit(val uint32) int {
	for i:=31; i>=0; i-- { // Start from left
		if getNthBit(val, uint32(i)) == 1 { // 0 index
			return i // Return first bit that has a 1 - zero-index
		}
	}
	return -1
}

func getNthBit(val, n uint32) int {
    // 1. reverse the golang endian
    nthBit := n
    // 2. move the nth bit to the first position
    movedVal := val >> nthBit
    // 3. mask the value, selecting only this first bit
    maskedValue := movedVal & 1
    return int(maskedValue)
    // can be shortened like so
    // return (val >> (32-n)) & 1
}

func calculateFromBinary(value, n int) int {
	// Calculate 2 to the power of n
	return value * int(math.Pow(float64(2), float64(n)))
}


func main() {
	input := []uint32{137}
	fmt.Println(input) // Print decimal value
	fmt.Println(fmt.Sprintf("%08b", input)) // Print binary value

	fmt.Println(getMax1Bit(input[0])) // 7

	var num int

	for i:=0; i<=7; i++ {
		fmt.Println(i, ":", getNthBit(input[0], uint32(i))) // zero-index
		num += calculateFromBinary(getNthBit(input[0], uint32(i)), i)
		fmt.Println("num is:", num)
	}

	

	// result := make([]uint32, 0)



	


}