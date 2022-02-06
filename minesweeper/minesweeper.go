package minesweeper

import (
	"fmt"
	"time"
)

// Returns a count of stars surrounding the given cell
func (b Board) GetSurroundStars(r, c int) int {
	var count int
	// Check rows
	for col := c - 1; col <= c+1; col++ {
		if b[r-1][col] == '*' { // Previous row
			count++
		}
		if b[r+1][col] == '*' { // Next row
			count++
		}
	}
	// Left
	if b[r][c-1] == '*' {
		count++
	}
	// Right
	if b[r][c+1] == '*' {
		count++
	}
	return count
}

func (b Board) Count() error {
	// Dimensions
	rows := len(b)    // Includes 2 extras
	cols := len(b[0]) // Inclues 2 extras
	fmt.Println("rows:", rows, "cols:", cols)

	for r := 1; r < rows-1; r++ { // Loop over rows
		for c := 1; c < cols-1; c++ { // Loop over cols in the row
			if string(b[r][c]) == "*" {
				fmt.Printf("%s", string(b[r][c]))
			} else {
				if b.GetSurroundStars(r, c) != 0 {
					fmt.Printf("%d", b.GetSurroundStars(r, c))
				} else {
					fmt.Printf(" ")
				}
			}
		}
		fmt.Printf("\n")
	}

	// Sleep for 5 seconds
	time.Sleep(5 * time.Second)

	return nil
}
