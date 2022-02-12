package wordsearch

import "fmt"

type board struct {
	rows  int
	cols  int
	words []string
}

// top -> bottom
func (b *board) top2bottom() []string {
	words := make([]string, b.rows)
	for c := 0; c < b.cols; c++ {
		for r := 0; r < b.rows; r++ {
			words[c] += string(b.words[r][c])
		}
	}
	return words
}

// bottom -> top
func (b *board) bottom2top() []string {
	words := make([]string, b.rows)
	for c := 0; c < b.cols; c++ {
		for r := b.rows - 1; r >= 0; r-- {
			words[c] += string(b.words[r][c])
		}
	}
	return words
}

// left -> right
func (b *board) left2right() []string {
	words := make([]string, b.rows)
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.cols; c++ {
			words[r] += string(b.words[r][c])
		}
	}
	return words
}

// right -> left
func (b *board) right2left() []string {
	words := make([]string, b.rows)
	for r := 0; r < b.rows; r++ {
		for c := b.cols - 1; c >= 0; c-- {
			words[r] += string(b.words[r][c])
		}
	}
	return words
}

func looptillbound(r, c int, b *board, words []string, i *int) {
	var diag_r, diag_c int = r, c              // Reset
	for diag_r != b.rows && diag_c != b.cols { // While not out of bounds
		words[*i] += string(b.words[diag_r][diag_c])
		diag_r++
		diag_c++
	}
	*i++
}

// top-left -> bottom-right
func (b *board) topleft2bottomright() []string {
	words := make([]string, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0

	// Sweep right across columns
	var r int = 0
	for c := 0; c < b.cols; c++ {
		var diag_r, diag_c int = r, c              // Reset
		for diag_r != b.rows && diag_c != b.cols { // While not out of bounds
			words[i] += string(b.words[diag_r][diag_c])
			diag_r++
			diag_c++
		}
		i++
	}
	// Sweep down across rows
	var c int = 0
	for r := 1; r < b.rows; r++ {
		var diag_r, diag_c int = r, c              // Reset
		for diag_r != b.rows && diag_c != b.cols { // While not out of bounds
			words[i] += string(b.words[diag_r][diag_c])
			diag_r++
			diag_c++
		}
		i++
	}

	return words
}

// bottom-right -> top-left

// top-right -> bottom-left

// bottom-left -> top-right

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	fmt.Println("puzzle:", puzzle)
	fmt.Println("words:", words)

	coords := make(map[string][2][2]int)
	var b board

	b.rows = len(puzzle)    // Length of puzzle is number of rows in the board
	b.cols = len(puzzle[0]) // Length of first row is number of columns in the board

	// Error handling
	for r := range puzzle { // Puzzle string length must equal to be rectangular
		if len(puzzle[r]) != b.cols {
			return nil, fmt.Errorf("puzzle is not rectangular")
		}
	}
	// Fill up words
	b.words = puzzle

	// Test
	fmt.Println("Top 2 bottom: ", (b.top2bottom()))
	fmt.Println("Bottom 2 top: ", (b.bottom2top()))
	fmt.Println("Left 2 right: ", (b.left2right()))
	fmt.Println("Right 2 left: ", (b.right2left()))
	fmt.Println("Top Left 2 Bottom Right: ", (b.topleft2bottomright()))

	return coords, nil
}
