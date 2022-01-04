package railfence

// Creates ZigZag pattern in a slice used for decoding
func createZigZag(rails int, msgLength int) [][]string {
	rail_slice := make([][]string, 0)

	// initialize and append rails of appropriate lengths
	for i := 0; i < rails; i++ {
		rail := make([]string, msgLength)
		rail_slice = append(rail_slice, rail)
	}

	var forward bool = true // Tracking direction
	var i int               // Tracking rail number

	for n := 0; n < msgLength; n++ {
		rail_slice[i][n] = "?"
		railZigger(&i, &forward, rails)
	}
	return rail_slice
}

// railZigger implements the novel rail switching logic
// Is a function to avoid code duplication
func railZigger(i *int, forward *bool, rails int) {
	// Shift i
	if *forward {
		*i++
	} else {
		*i--
	}
	// Change direction at end of rails
	if *i == rails-1 {
		*forward = false
	} else if *i == 0 {
		*forward = true
	}
}

func Encode(message string, rails int) string {
	var result string                   // Return value
	var i int                           // Tracking rail number
	var forward bool = true             // Tracking direction
	rail_slice := make([]string, rails) // Each entry of slice represents number of rails
	for _, c := range message {         // Loop per character
		rail_slice[i] += string(c) // Append to rail
		railZigger(&i, &forward, rails)
	}
	for _, c := range rail_slice { // Form message
		result += c
	}
	return result
}

func Decode(message string, rails int) string {
	rail_slice := createZigZag(rails, len(message))

	// // Replace "?" per rail with message
	var i int // Message index tracker
	for _, rail := range rail_slice {
		// Loop over rail
		for n, c := range rail {
			if c == string('?') {
				rail[n] = string(message[i])
				i++
			}
		}
	}
	// Get back the string
	var result string       // Return value
	i = 0                   // Tracking rail number
	var forward bool = true // Tracking direction
	for x := 0; x < len(message); x++ {
		result += rail_slice[i][x]
		railZigger(&i, &forward, rails)
	}
	return result
}
