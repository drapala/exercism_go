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

type coordinates struct {
	start []int
	end   []int
}

type wordmap struct {
	word  string
	coord [][]int
}

// top -> bottom
func (b *board) top2bottom() []wordmap {
	wm := make([]wordmap, b.rows)
	for c := 0; c < b.cols; c++ {
		for r := 0; r < b.rows; r++ {
			wm[c].word += string(b.words[r][c])
			wm[c].coord.start = []int{c, 0}
			wm[c].coord.end = []int{c, b.rows - 1}
		}
	}
	return wm
}

// bottom -> top
func (b *board) bottom2top() []wordmap {
	wm := make([]wordmap, b.rows)
	for c := 0; c < b.cols; c++ {
		for r := b.rows - 1; r >= 0; r-- {
			wm[c].word += string(b.words[r][c])
			wm[c].coord.start = []int{c, b.rows - 1}
			wm[c].coord.end = []int{c, 0}
		}
	}
	return wm
}

// left -> right
func (b *board) left2right() []wordmap {
	wm := make([]wordmap, b.rows)
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.cols; c++ {
			wm[r].word += string(b.words[r][c])
			wm[r].coord.start = []int{0, r}
			wm[r].coord.end = []int{b.cols - 1, r}
		}
	}
	return wm
}

// right -> left
func (b *board) right2left() []wordmap {
	wm := make([]wordmap, b.rows)
	for r := 0; r < b.rows; r++ {
		for c := b.cols - 1; c >= 0; c-- {
			wm[r].word += string(b.words[r][c])
			wm[r].coord.start = []int{b.cols - 1, r}
			wm[r].coord.end = []int{0, r}
		}
	}
	return wm
}

// top-left -> bottom-right
func (b *board) topleft2bottomright() []wordmap {
	wm := make([]wordmap, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep right across columns
	var r int = 0 // Pin it
	for c := 0; c < b.cols; c++ {
		loop_topleft2bottomright(r, c, b, wm, &i)
	}
	// Sweep down across rows
	var c int = 0                 // Pin it
	for r := 1; r < b.rows; r++ { // Note this is shifted by 1 to not double count
		loop_topleft2bottomright(r, c, b, wm, &i)
	}
	return wm
}

// bottom-right -> top-left
func (b *board) bottomright2topleft() []wordmap {
	wm := make([]wordmap, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep left across columns
	var r int = b.rows - 1 // Pin it
	for c := b.cols - 1; c >= 0; c-- {
		loop_bottomright2topleft(r, c, b, wm, &i)
	}
	// Sweep up across rows
	var c int = b.cols - 1             // Pin it
	for r := b.rows - 2; r >= 0; r-- { // Note this is shifted by 1 to not double count
		loop_bottomright2topleft(r, c, b, wm, &i)
	}
	return wm
}

// top-right -> bottom-left
func (b *board) topright2bottomleft() []wordmap {
	wm := make([]wordmap, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep left across columns
	var r int = 0 // Pin it
	for c := b.cols - 1; c >= 0; c-- {
		loop_topright2bottomleft(r, c, b, wm, &i)
	}
	// Sweep down across rows
	var c int = b.cols - 1        // Pin it
	for r := 1; r < b.rows; r++ { // Note this is shifted by 1 to not double count
		loop_topright2bottomleft(r, c, b, wm, &i)
	}
	return wm
}

// bottom-left -> top-right
func (b *board) bottomleft2topright() []wordmap {
	wm := make([]wordmap, b.cols+b.rows-1) // -1 because we don't count [0,0] twice
	var i int = 0
	// Sweep right across columns
	var r int = b.rows - 1 // Pin it
	for c := 0; c < b.cols; c++ {
		loop_bottomleft2topright(r, c, b, wm, &i)
	}
	// Sweep up across rows
	var c int = 0                      // Pin it
	for r := b.rows - 2; r >= 0; r-- { // Note this is shifted by 1 to not double count
		loop_bottomleft2topright(r, c, b, wm, &i)
	}
	return wm
}

// Helpers
func loop_topleft2bottomright(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord.start = []int{c, r}
		wm[*i].coord.end = []int{diag_c, diag_r}

		diag_r++
		diag_c++
	}
	*i++
}

func loop_bottomright2topleft(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord.start = []int{c, r}
		wm[*i].coord.end = []int{diag_c, diag_r}

		diag_r--
		diag_c--
	}
	*i++
}

func loop_topright2bottomleft(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord.start = []int{c, r}
		wm[*i].coord.end = []int{diag_c, diag_r}
		diag_r++
		diag_c--
	}
	*i++
}

func loop_bottomleft2topright(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord.start = []int{c, r}
		wm[*i].coord.end = []int{diag_c, diag_r}
		diag_r--
		diag_c++
	}
	*i++
}

func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	fmt.Println("puzzle:", puzzle)
	fmt.Println("words:", words)
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
	fmt.Println("############################################################")
	fmt.Println("#0 | Top 2 bottom: ", (b.top2bottom()))
	fmt.Println("############################################################")
	fmt.Println("#1 | Bottom 2 top: ", (b.bottom2top()))
	fmt.Println("############################################################")
	fmt.Println("#2 | Left 2 right: ", (b.left2right()))
	fmt.Println("############################################################")
	fmt.Println("#3 | Right 2 left: ", (b.right2left()))
	fmt.Println("############################################################")
	fmt.Println("#4 | Top Left 2 Bottom Right: ", (b.topleft2bottomright()))
	fmt.Println("############################################################")
	fmt.Println("#5 | Bottom Right 2 Top Left: ", (b.bottomright2topleft()))
	fmt.Println("############################################################")
	fmt.Println("#6 | Top Right 2 Bottom Left: ", (b.topright2bottomleft()))
	fmt.Println("############################################################")
	fmt.Println("#7 | Bottom Left 2 Top Right: ", (b.bottomleft2topright()))
	fmt.Println("############################################################")

	searchlist := [][]wordmap{b.top2bottom(), b.bottom2top(), b.left2right(), b.right2left(), b.topleft2bottomright(), b.bottomright2topleft(), b.topright2bottomleft(), b.bottomleft2topright()}

	// Coordinates to return
	coords := make(map[string][2][2]int)

	// Identify hits
	for _, word := range words {
		for s, search := range searchlist {
			index, shift := findwordinslice(word, search)
			if index != -1 {
				fmt.Println("Found:", word, " in search slice #: ", s, " at word #: ", index, " in ", search[index], " starting at index #: ", shift)

				// Add to coords
				coords[word] = search[index].coord

			}
		}
	}

	return coords, nil
}

// Uses Regex to check if word is contained in any of the words
func findwordinslice(word string, search []wordmap) (int, int) {
	for i, w := range search {
		if strings.Contains(w.word, word) {
			return i, strings.Index(w.word, word) // Returns index of word in search slice
		}
	}
	return -1, -1
}
