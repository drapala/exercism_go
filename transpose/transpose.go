package transpose

func Transpose(input []string) []string {
	// Generate dimensions based on longest string
	rows := len(input)
	var cols int
	for _, c := range input {
		if len(c) > cols {
			cols = len(c)
		}
	}

	// Output slice
	var output []string = make([]string, 0)

	// Loop through each column and go right on each row
	// If no character found, if upcoming rows in that column have characters - check downward and pad with space
	var word string
	var hasChar bool

	for c := 0; c < cols; c++ { // Loop through each column
		for r := 0; r < rows; r++ { // Loop through each row
			if c > len(input[r])-1 { // If we are past the end of the column for this row
				// Check if any of the next rows have a valid character
				for i := r + 1; i < rows; i++ {
					if len(input[i]) > c {
						hasChar = true
						break
					}
				}
				// If hasChar, pad with space
				if hasChar {
					word += " "
					hasChar = false // Reset
				}
			} else { // Not past the end of the row - add character as is
				word += string(input[r][c])
			}
		}
		output = append(output, word)
		word = ""
	}
	return output
}
