package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Define the Matrix type here.
type Matrix struct {
	rows     [][]int
	num_rows int
	num_cols int
}

func New(s string) (*Matrix, error) {
	// Matrix to return
	var m = Matrix{}
	// Split by newline to get rows
	rows := strings.Split(s, "\n")
	// To store each individual element per row
	var elements []string
	// Split each row by space to get individual elements
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
	// Return pointer to matrix
	return &m, nil
}

// Cols and Rows must return the results without affecting the matrix.
func (m *Matrix) Rows() [][]int {
	return m.rows
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
	if row < 0 || row >= m.num_rows-1 || col < 0 || col >= m.num_cols-1 {
		return false
	} else {
		m.rows[row][col] = val
		return true
	}
}

func main() {
	var s = "1 2 3\n4 5 6\n7 8 9\n 8 7 6"
	m, _ := New(s)
	fmt.Println("rows", m.Rows())
	fmt.Println("cols", m.Cols())

	m.Set(1, 1, 99)
	fmt.Println("rows", m.Rows())
	fmt.Println("cols", m.Cols())
}
