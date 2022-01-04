package railfence

import "fmt"

func Encode(message string, rails int) string {
	var result string                   // Return value
	var i int                           // Tracking rail number
	var forward bool = true             // Tracking direction
	rail_slice := make([]string, rails) // Each entry of slice represents number of rails
	for _, c := range message {         // Loop per character
		rail_slice[i] += string(c) // Append to rail
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
	for _, c := range rail_slice { // Form message
		result += c
	}
	return result
}

func insertSpaces (in string, num int) string {
	for i := 1; i <= num; i++ {
		in += " "
	}
	return in
}

func Decode(message string, rails int) string {
	// Determine spaces per rail
	spaces := make([]int, rails-1)
	var railspace int = 1
	for i := 0; i < len(spaces); i++ {
		spaces[i] = railspace
		railspace += 2
	}
	fmt.Println("spaces:", spaces)

	rail_slice := make([]string, rails)

	// Insert leading spaces per rail
	for i := 0; i < rails; i++ {
		rail_slice[i] = insertSpaces(rail_slice[i], i)
		//rail_slice[i] += "_" // Remove!
	}
	
	var r int // For tracking rail number

	for i := 0; i < len(message); i++{ // Loop over messages to insert
		rail_slice[r] += string(message[i])
		
		// Insert spaces appropriately to rail as per "spaces"
		switch r {
		case 0:
			rail_slice[r] = insertSpaces(rail_slice[r], spaces[rails-2])
		case rails-1:
			rail_slice[r] = insertSpaces(rail_slice[r], spaces[rails-2])
		default:
			rail_slice[r] = insertSpaces(rail_slice[r], spaces[(rails - 2) - r])
		}
		// If rail is == message length, increment railnum
		if len(rail_slice[r]) >= len(message) {
			r++
		}
	}
	// Print railslice
	for _, c := range(rail_slice) {
		fmt.Println(c)
	}



	// Form string to return
	var result string
	r = 0
	var forward bool = true             // Tracking direction
	// Gather decoded string
	for i:=0; i < len(message); i++ {
		result += string(rail_slice[r][i])
		// Shift i
		if forward {
			r++
		} else {
			r--
		}
		// Change direction at end of rails
		if r == rails-1 {
			forward = false
		} else if r == 0 {
			forward = true
		}
	}

	fmt.Println("result:", result)

	return result
}
