package diamond

import (
	"fmt"
)

type board struct {
	rows      string // Finished board
	dimension int    // Length, Height
	char      byte   // Char to get to before flipping
}

func needToInsert(pos, i, row, dimension int, flip bool) bool {
	if !flip { // Up until halfway point
		if pos == (i-(row-1)) || pos == (i+(row-1)) {
			return true
		}
		return false
	} else { // Past the halfway point
		if pos == (i+row-dimension) || pos == (i-row+dimension) {
			return true
		}
	}
	return false // Default
}

func (b *board) generateBoard() {
	b.rows = ""                 // To contain finished board
	i := int(b.char) - int('A') // Distance from A

	var ascii int         // Ascii of current row
	var flip bool = false // Reached halfway point
	var half int          // The halfway point

	for row := 1; row <= b.dimension; row++ { // Loop through rows
		for pos := 0; pos < b.dimension; pos++ { // Loop through positions
			if needToInsert(pos, i, row, b.dimension, flip) { // If we need to insert character
				if flip { // If we are past halfway point
					ascii = int(b.char) - (row - half)
					b.rows += string(rune(ascii))
				} else {
					ascii = int('A') + row - 1
					b.rows += string(rune(ascii))
				}
				if ascii == int(b.char) && pos == (i+(row-1)) { // Mark Halfway point after second write
					flip = true
					half = row
				}
			} else {
				b.rows += " " // Change to space after
			}
		}
		// After looping through the positions
		b.rows += "\n"
	}
}

func Gen(char byte) (string, error) {
	// Distance from A
	input_num := int(char) - int('A')

	// Error handling
	if input_num < 0 || input_num > 25 { // Outside of A - Z
		return "", fmt.Errorf("invalid input: %s", string(char))
	}

	// Generate empty board
	my_board := board{}
	my_board.dimension = input_num*2 + 1
	my_board.char = char

	// Fill board
	my_board.generateBoard()

	return my_board.rows, nil
}
