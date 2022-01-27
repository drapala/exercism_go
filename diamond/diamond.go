package diamond

import (
	"fmt"
	"time"
)

type board struct {
	rows      string // Finished board
	dimension int    // Length, Height
	char      byte   // Char to get to before flipping
}

func needToInsert(pos, i, row int, flip bool) bool {
	fmt.Println("flip?", flip)
	if !flip { // Up until halfway point
		if pos == (i-(row-1)) || pos == (i+(row-1)) {
			return true
		}
		return false
	} else { // Past the halfway point
		// Continue from here
	}
	return false // Default
}

func (b board) generateBoard() {
	b.rows = ""                 // To contain finished board
	i := int(b.char) - int('A') // Distance from A
	fmt.Println("Distance from A: ", i)
	fmt.Println("Board dimension: ", b.dimension)
	fmt.Println("Halfway at:", string(rune(int(b.char))))

	var ascii int         // Ascii of current row
	var flip bool = false // Reached halfway point
	var half int          // The halfway point

	for row := 1; row <= b.dimension; row++ { // Loop through rows
		for pos := 0; pos < b.dimension; pos++ { // Loop through positions
			if needToInsert(pos, i, row, flip) { // If we need to insert character
				if flip { // If we are past halfway point
					ascii = int(b.char) - (row - half)
				} else {
					ascii = int('A') + row - 1
					b.rows += string(rune(ascii))
				}
				if ascii == int(b.char) && pos == (i+(row-1)) { // Halfway point after second write
					flip = true
					half = row
				}
			} else {
				b.rows += "." // Change to space after
			}
		}
		// After looping through the positions
		b.rows += "\n"
		fmt.Println(b.rows)
		time.Sleep(500 * time.Millisecond)
	}
	// Go to sleep
	// time.Sleep(10 * time.Second)
}

func Gen(char byte) (string, error) {

	fmt.Println("input:", string(char))
	input_num := int(char) - int('A') // Distance from A

	// Generate empty board
	my_board := board{}
	my_board.dimension = input_num*2 + 1
	my_board.char = char

	// Fill board
	my_board.generateBoard()

	// Error handle outsize Z

	return "", nil
}
