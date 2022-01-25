package diamond

import (
	"fmt"
)

type board struct {
	rows      string
	dimension int
	char      byte
}

func (b board) generateBoard() {
	b.rows = ""

	i := int(b.char) - int('A') // Distance from A
	fmt.Println("Distance from A: ", i)
	fmt.Println("Board dimension: ", b.dimension)

	curr_char := rune('A') // Rune of current row

	for row := 1; row <= b.dimension; row++ { // Loop through rows
		for pos := 0; pos < b.dimension; pos++ { // Loop through positions
			if pos == (i-1-(row-1)) || pos == (i-1+(row-1)) {
				b.rows += string('A' + row - 1) // Need to fix
			} else {
				b.rows += " "
			}
			curr_char++ // Next letter
		}
		b.rows += "\n"
		fmt.Println(b.rows)
	}
}

func Gen(char byte) (string, error) {

	fmt.Println("input:", string(char))
	input_num := int(char) - int('A') // Distance from A

	// Generate empty board
	my_board := board{}
	my_board.dimension = input_num*2 - 1
	my_board.char = char

	my_board.generateBoard()

	// Error handle outsize Z

	return "", nil
}
