package minesweeper

import (
	"fmt"
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

// Check if a board is clean or not
func (b Board) CheckClean() error {
	// Dimension check
	rows := len(b)
	cols := len(b[0])
	// Loop over rows
	for r := 0; r < rows; r++ {
		// Check if same size as first row
		if len(b[r]) != cols {
			return fmt.Errorf("row %d has %d columns, expected %d", r, len(b[r]), cols)
		}
		// Edge '|' check
		if (b[r][0] != '|') && (b[r][0] != '+') { // Left side
			return fmt.Errorf("left edge of board has %c, expected | or +", b[r][0])
		}
		if (b[r][cols-1] != '|') && (b[r][cols-1] != '+') { // Right side
			return fmt.Errorf("right edge of board has %c, expected | or +", b[r][cols-1])
		}
		// Foreign character check
		for c := 0; c < cols; c++ {
			// Must be *, +, -, | or ' '
			d := b[r][c]
			if (d != '*') && (d != '+') && (d != '-') && (d != '|') && (d != ' ') {
				return fmt.Errorf("board has %c at row %d, col %d, expected *, |, +, or -", d, r, c)
			}
		}
	}
	// Return nil if all checks pass
	return nil
}

func (b Board) Count() error {
	// Error check
	err := b.CheckClean()
	if err != nil {
		return err
	}
	// Dimensions
	rows := len(b)    // Includes 2 extras
	cols := len(b[0]) // Inclues 2 extras
	// Loop over rows
	for r := 1; r < rows-1; r++ {
		for c := 1; c < cols-1; c++ { // Loop over cols in the row
			if string(b[r][c]) != "*" { // Don't overwrite stars
				if b.GetSurroundStars(r, c) != 0 {
					count := byte(b.GetSurroundStars(r, c))
					b[r][c] = byte(count + 48) // Shift to ASCII
				}
			}
		}
	}
	return nil
}
