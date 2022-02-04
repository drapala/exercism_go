package queenattack

import (
	"errors"
	"math"
)

// Create map of letters to column number
var col_map = map[string]int{
	"a": 0,
	"b": 1,
	"c": 2,
	"d": 3,
	"e": 4,
	"f": 5,
	"g": 6,
	"h": 7,
}

var row_map = map[string]int{
	"8": 0,
	"7": 1,
	"6": 2,
	"5": 3,
	"4": 4,
	"3": 5,
	"2": 6,
	"1": 7,
}

func pos2array(pos string) (int, int, error) {
	// Error handling
	if len(pos) != 2 { // Not 2 characters
		return 0, 0, errors.New("invalid length of position")
	}
	// 2 characters from this point on
	firstChar := pos[0]
	secondChar := pos[1]
	// Error handling
	if firstChar < 'a' || firstChar > 'h' { // First char not in range
		return 0, 0, errors.New("invalid first character of position")
	} else if secondChar < '1' || secondChar > '8' { // Second char not in range
		return 0, 0, errors.New("invalid second character of position")
	}
	// Get index mapping
	row, col := col_map[string(firstChar)], row_map[string(secondChar)]
	return row, col, nil
}

func CheckIfDiagonal(w_r, w_c, b_r, b_c int) bool {
	// Check if on same diagonal
	return math.Abs(float64(w_r-b_r)) == math.Abs(float64(w_c-b_c))
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	w_r, w_c, w_err := pos2array(whitePosition)
	b_r, b_c, b_err := pos2array(blackPosition)
	// Error handling
	if w_err != nil {
		return false, w_err
	} else if b_err != nil {
		return false, b_err
	} else if whitePosition == blackPosition { // Same position, no attack
		return false, errors.New("same position")
	}
	// Same row or column
	if w_c == b_c { // Check if on same column
		return true, nil
	} else if w_r == b_r { // Check if on same row
		return true, nil
	}
	// Check if on same diagonal
	return CheckIfDiagonal(w_r, w_c, b_r, b_c), nil
}
