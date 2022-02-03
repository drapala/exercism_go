package matrix

import (
	"errors"
	"strconv"
	"strings"
)

// Matrix is copied from matrix.go at the end of this file.

// Pair:
// Contains row and column co-ordinates containing saddle points
type Pair struct {
	row int
	col int
}

func (m *Matrix) Saddle() []Pair {
	pair := make([]Pair, 0)
	// Start assuming true, will flip to false otherwise
	saddle_row := true
	saddle_col := true
	for r, row := range m.Rows() { // Loop over rows
		for c, item := range row { // Loop over cols in the row
			// Check if row is a saddle point
			for _, v := range m.Rows()[r] {
				if item < v { // Not a saddle point since item is lt other vals in row
					saddle_row = false
				}
			}
			// Check if col is a saddle point
			for _, v := range m.Cols()[c] {
				if item > v { // Not a saddle point since item is gt other vals in row
					saddle_col = false
				}
			}
			// Append if both still true
			if saddle_row && saddle_col {
				pair = append(pair, Pair{r, c})
			}
			// Reset flag
			saddle_row = true
			saddle_col = true
		}
	}
	return pair
}

// #############################################################################
// 								COPIED FROM MATRIX.GO
// #############################################################################

// Define the Matrix type here.
type Matrix struct {
	rows     [][]int
	num_rows int
	num_cols int
}

// Check dimensions of all rows and cols are same
func SameDimensions(m *Matrix) bool {
	rows := m.Rows()
	cols := m.Cols()

	// Loop over rows and check dimensions
	for i := 0; i < len(rows); i++ {
		if len(rows[i]) != len(rows[0]) {
			return false
		}
	}
	// Loop over cols and check dimensions
	for i := 0; i < len(cols); i++ {
		if len(cols[i]) != len(cols[0]) {
			return false
		}
	}
	// All rows and cols have same dimensions
	return true
}

// Create a new matrix from a string
func New(s string) (*Matrix, error) {
	// Matrix to return
	var m = Matrix{}
	// Split by newline to get rows slice
	rows := strings.Split(s, "\n")
	// To store each individual element per row
	var elements []string
	// Loop over rows slice
	for _, row := range rows {
		// Trim leading and trailing spaces
		row = strings.TrimSpace(row)
		// Split row into elements by splitting by space
		elements = strings.Split(row, " ")
		// Clean up each element and convert to integer, append to the row slice
		new_row := []int{}
		for _, element := range elements {
			n, err := strconv.Atoi(strings.TrimSpace(element))
			// If element cannot be converted to integer, return error
			if err != nil {
				return nil, err
			}
			// Append element to row slice
			new_row = append(new_row, n)
		}
		// Append row slice to matrix
		m.rows = append(m.rows, new_row)
	}
	// Matrix dimensions
	m.num_cols = len(m.rows[0])
	m.num_rows = len(m.rows)
	// Check dimensions
	if !SameDimensions(&m) {
		return nil, errors.New("Matrix dimensions do not match")
	}
	// Return pointer to matrix
	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Rows() [][]int {
	// Because we want to return an independent copy of the slice
	var rows [][]int
	var new_row []int
	// Loop over rows
	for _, row := range m.rows {
		new_row = make([]int, len(row))
		// Copy values
		copy(new_row, row)
		// Append to return
		rows = append(rows, new_row)
	}
	return rows
}

// Loop over rows and return cols
func (m *Matrix) Cols() [][]int {
	var cols [][]int
	// Loop over num_cols and create columns
	for i := 0; i < m.num_cols; i++ {
		new_col := []int{}
		for j := 0; j < m.num_rows; j++ {
			// Append element to col slice
			new_col = append(new_col, m.rows[j][i])
		}
		// Append col slice to matrix
		cols = append(cols, new_col)
	}
	return cols
}

func (m *Matrix) Set(row, col, val int) bool {
	// Out of range
	if row < 0 || row > m.num_rows-1 || col < 0 || col > m.num_cols-1 {
		return false
	} else {
		m.rows[row][col] = val
		return true
	}
}
