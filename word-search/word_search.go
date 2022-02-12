package wordsearch

import (
	"fmt"
	"strings"
)

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

// top-left -> bottom-right
func (b *board) topleft2bottomright() []string {
	words := make([]string, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep right across columns
	var r int = 0 // Pin it
	for c := 0; c < b.cols; c++ {
		loop_topleft2bottomright(r, c, b, words, &i)
	}
	// Sweep down across rows
	var c int = 0                 // Pin it
	for r := 1; r < b.rows; r++ { // Note this is shifted by 1 to not double count
		loop_topleft2bottomright(r, c, b, words, &i)
	}
	return words
}

// bottom-right -> top-left
func (b *board) bottomright2topleft() []string {
	words := make([]string, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep left across columns
	var r int = b.rows - 1 // Pin it
	for c := b.cols - 1; c >= 0; c-- {
		loop_bottomright2topleft(r, c, b, words, &i)
	}
	// Sweep up across rows
	var c int = b.cols - 1             // Pin it
	for r := b.rows - 2; r >= 0; r-- { // Note this is shifted by 1 to not double count
		loop_bottomright2topleft(r, c, b, words, &i)
	}
	return words
}

// top-right -> bottom-left
func (b *board) topright2bottomleft() []string {
	words := make([]string, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep left across columns
	var r int = 0 // Pin it
	for c := b.cols - 1; c >= 0; c-- {
		loop_topright2bottomleft(r, c, b, words, &i)
	}
	// Sweep down across rows
	var c int = b.cols - 1        // Pin it
	for r := 1; r < b.rows; r++ { // Note this is shifted by 1 to not double count
		loop_topright2bottomleft(r, c, b, words, &i)
	}
	return words
}

// bottom-left -> top-right
func (b *board) bottomleft2topright() []string {
	words := make([]string, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep right across columns
	var r int = b.rows - 1 // Pin it
	for c := 0; c < b.cols; c++ {
		loop_bottomleft2topright(r, c, b, words, &i)
	}
	// Sweep up across rows
	var c int = 0                      // Pin it
	for r := b.rows - 2; r >= 0; r-- { // Note this is shifted by 1 to not double count
		loop_bottomleft2topright(r, c, b, words, &i)
	}
	return words
}

// Helpers
func loop_topleft2bottomright(r, c int, b *board, words []string, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		words[*i] += string(b.words[diag_r][diag_c])
		diag_r++
		diag_c++
	}
	*i++
}

func loop_bottomright2topleft(r, c int, b *board, words []string, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		words[*i] += string(b.words[diag_r][diag_c])
		diag_r--
		diag_c--
	}
	*i++
}

func loop_topright2bottomleft(r, c int, b *board, words []string, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		words[*i] += string(b.words[diag_r][diag_c])
		diag_r++
		diag_c--
	}
	*i++
}

func loop_bottomleft2topright(r, c int, b *board, words []string, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		words[*i] += string(b.words[diag_r][diag_c])
		diag_r--
		diag_c++
	}
	*i++
}

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

	// Print
	fmt.Println("#0 | Top 2 bottom: ", (b.top2bottom()))
	fmt.Println("#1 | Bottom 2 top: ", (b.bottom2top()))
	fmt.Println("#2 | Left 2 right: ", (b.left2right()))
	fmt.Println("#3 | Right 2 left: ", (b.right2left()))
	fmt.Println("#4 | Top Left 2 Bottom Right: ", (b.topleft2bottomright()))
	fmt.Println("#5 | Bottom Right 2 Top Left: ", (b.bottomright2topleft()))
	fmt.Println("#6 | Top Right 2 Bottom Left: ", (b.topright2bottomleft()))
	fmt.Println("#7 | Bottom Left 2 Top Right: ", (b.bottomleft2topright()))

	searchlist := [][]string{b.top2bottom(), b.bottom2top(), b.left2right(), b.right2left(), b.topleft2bottomright(), b.bottomright2topleft(), b.topright2bottomleft(), b.bottomleft2topright()}

	// Look for hits
	for _, word := range words {
		for s, search := range searchlist {
			if findwordinslice(word, search) != -1 {
				fmt.Println("Found:", word, " in search slice #: ", s, " at word #: ", findwordinslice(word, search), " in ", search[findwordinslice(word, search)])
			}
		}
	}

	return coords, nil
}

// Uses Regex to check if word is contained in any of the words
func findwordinslice(word string, search []string) int {
	for i, w := range search {
		if strings.Contains(w, word) {
			return i // Returns index of word in search slice
		}
	}
	return -1
}
