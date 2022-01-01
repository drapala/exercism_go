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

// Reverse a slice
func reverseSlice(input []byte) []byte {
	output := make([]byte, 0)
	for i:=len(input)-1; i>=0; i-- {
		output = append(output, input[i])
	}
	return output
}

func main() {
	input := []uint32{0x4000}
	fmt.Println(input) // Print decimal value
	fmt.Println(fmt.Sprintf("%08b", input)) // Print binary value
	// 1 0000000 0000000

	seven_bit := make([]byte, 0)
	var c, nth, num int

	for i:=0; i<=getMax1Bit(input[0]); i++ {
		nth = getNthBit(input[0], uint32(i))
		num += calculateFromBinary(nth, c)
		c++
		// 7th bit, reset
		if c == 7 { 
			seven_bit = append(seven_bit, byte(num))
			c = 0
			num = 0
		} 
		// Last bit - append to slice
		if i == getMax1Bit(input[0]){ 
			seven_bit = append(seven_bit, byte(num))
		}
	}
	seven_bit = reverseSlice(seven_bit)

	// Add MSB
	for i:=0; i<len(seven_bit); i++ {
		if i != len(seven_bit)-1 { // Not last bit - set MSB = 1
			seven_bit[i] += byte(calculateFromBinary(1, 7))
		}
		fmt.Println(fmt.Sprintf("%x", seven_bit[i]))
	}
	// Print hex value
	fmt.Println(fmt.Sprintf("%x", seven_bit)) // Print hex value
}