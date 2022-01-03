package main

import (
	"fmt"
	"math"
)

// Starts from the left of a 32-bit integer and returns the first bit that has a 1.
func getMax1Bit(val uint32) int {
	for i:=31; i>=0; i-- { // Start from left
		if getNthBit(val, uint32(i)) == 1 { // 0 index
			return i // Return first bit that has a 1 - zero-index
		}
	}
	return -1
}

// Returns the value of the nth bit of a 32-bit integer.
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

// Convert a binary number to decimal for a given power of 2
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

// Convets a single unsigned 32 bit integer to VLQ representation
func convertint32ToVLQ (input uint32) []byte {
	// Handle zero case since we operate based on first non-zero nth byte here
	if input == 0 {
		return []byte{0x0}
	} 

	vlq_bit := make([]byte, 0) // Return slice
	var c, num_2, num_10 int // Intermediate variables
	// c: used to track the 0-7 moving index
	// num_2: value of the nth bit from binary conversion of input
	// num_10: decimal value of num

	// Loop from right to left - until the last bit that is 1
	for i:=0; i<=getMax1Bit(input); i++ {
		num_2 = getNthBit(input, uint32(i)) // Get the nth bit from binary
		num_10 += calculateFromBinary(num_2, c) // Convert and append to base 10
		c++ // Increment c to keep track
		// We've filled up the 7th bit, append and reset
		if c == 7 { 
			vlq_bit = append(vlq_bit, byte(num_10))
			c = 0
			num_10 = 0
		} 
		// Last bit - may not be a factor of 7, so append to slice
		// If c == 0, that means we just appended so no need to append again
		if i == getMax1Bit(input) && c != 0 { 
			vlq_bit = append(vlq_bit, byte(num_10))
		}
	}
	vlq_bit = reverseSlice(vlq_bit) // Reverse slice to match output format

	// Add MSB bit
	for i:=0; i<len(vlq_bit); i++ {
		if i != len(vlq_bit)-1 { // Not last bit - set MSB = 1
			vlq_bit[i] += byte(calculateFromBinary(1, 7))
		}
	}
	return vlq_bit
}

// Get the end of a VLQ sequence based on MSB = 0
func getVLQend(byte_input []byte, track_left int) int {
	for i:=track_left; i<len(byte_input); i++ {
		if getNthBit(uint32(byte_input[i]), 7) == 0 { // MSB is 0 - end of VLQ
			return i
		}
	}
	return -1 // Not a valid VLQ
}

func main() {
	// #######
	// Decode
	// #######
	byte_input := []byte{0x8f, 0xff, 0xff, 0xff, 0x7f}
	fmt.Printf("input: %x\n", byte_input)

	var track_left, track_right int
	var_int := make([]uint32, 0) // Return slice
	var num_2, num_10, pow int

	for true { // While loop, we break inside
		track_right = getVLQend(byte_input, track_left)
		if track_right == -1 { // Not a valid VLQ
			break
		}
		// Loop right to left within this VLQ
		for i:=track_right; i>=track_left; i-- {
			for c:=0; c<=6; c++ { // Leave out the MSB
				num_2 = getNthBit(uint32(byte_input[i]), uint32(c))
				num_10 += calculateFromBinary(num_2, pow)
				pow++ // Power of 2 increases right to left regardless of i or c
			}
		}
		var_int = append(var_int, uint32(num_10)) // Append VLQ to slice
		track_left = track_right + 1 // Move to the next VLQ
		num_2, num_10, pow = 0, 0, 0 // Reset power of 2 for next VLQ run
	}
	fmt.Printf("output: %x", var_int)
}