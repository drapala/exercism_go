package wordsearch

import (
	"errors"
	"strings"
)

type board struct {
	rows  int
	cols  int
	words []string
}

type wordmap struct {
	word  string
	coord [2][2]int
}

// top -> bottom
func (b *board) top2bottom() []wordmap {
	wm := make([]wordmap, b.cols)
	for c := 0; c < b.cols; c++ {
		for r := 0; r < b.rows; r++ {
			wm[c].word += string(b.words[r][c])
			wm[c].coord[0] = [2]int{c, 0}
			wm[c].coord[1] = [2]int{c, b.rows - 1}
		}
	}
	return wm
}

// bottom -> top
func (b *board) bottom2top() []wordmap {
	wm := make([]wordmap, b.cols)
	for c := 0; c < b.cols; c++ {
		for r := b.rows - 1; r >= 0; r-- {
			wm[c].word += string(b.words[r][c])
			wm[c].coord[0] = [2]int{c, b.rows - 1}
			wm[c].coord[1] = [2]int{c, 0}
		}
	}
	return wm
}

// left -> right
func (b *board) left2right() []wordmap {
	wm := make([]wordmap, b.cols)
	for r := 0; r < b.rows; r++ {
		for c := 0; c < b.cols; c++ {
			wm[r].word += string(b.words[r][c])
			wm[r].coord[0] = [2]int{0, r}
			wm[r].coord[1] = [2]int{b.cols - 1, r}
		}
	}
	return wm
}

// right -> left
func (b *board) right2left() []wordmap {
	wm := make([]wordmap, b.cols)
	for r := 0; r < b.rows; r++ {
		for c := b.cols - 1; c >= 0; c-- {
			wm[r].word += string(b.words[r][c])
			wm[r].coord[0] = [2]int{b.cols - 1, r}
			wm[r].coord[1] = [2]int{0, r}
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
		wm[*i].coord[0] = [2]int{c, r}
		wm[*i].coord[1] = [2]int{diag_c, diag_r}

		diag_r++
		diag_c++
	}
	*i++
}

func loop_bottomright2topleft(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord[0] = [2]int{c, r}
		wm[*i].coord[1] = [2]int{diag_c, diag_r}

		diag_r--
		diag_c--
	}
	*i++
}

func loop_topright2bottomleft(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord[0] = [2]int{c, r}
		wm[*i].coord[1] = [2]int{diag_c, diag_r}
		diag_r++
		diag_c--
	}
	*i++
}

func loop_bottomleft2topright(r, c int, b *board, wm []wordmap, i *int) {
	var diag_r, diag_c int = r, c                                              // Reset
	for diag_r != -1 && diag_r != b.rows && diag_c != -1 && diag_c != b.cols { // While not out of bounds
		wm[*i].word += string(b.words[diag_r][diag_c])
		wm[*i].coord[0] = [2]int{c, r}
		wm[*i].coord[1] = [2]int{diag_c, diag_r}
		diag_r--
		diag_c++
	}
	*i++
}

// Main function
func Solve(words []string, puzzle []string) (map[string][2][2]int, error) {
	var b board

	b.rows = len(puzzle)    // Length of puzzle is number of rows in the board
	b.cols = len(puzzle[0]) // Length of first row is number of columns in the board

	// Error handling
	for r := range puzzle { // Puzzle string length must equal to be rectangular
		if len(puzzle[r]) != b.cols {
			return nil, errors.New("puzzle is not rectangular")
		}
	}
	// Fill up words
	b.words = puzzle

	searchlist := [][]wordmap{b.top2bottom(), b.bottom2top(), b.left2right(), b.right2left(), b.topleft2bottomright(), b.bottomright2topleft(), b.topright2bottomleft(), b.bottomleft2topright()}

	// Coordinates to return
	coords := make(map[string][2][2]int)

	var found bool = false // if found is true, we can stop searching

	// Identify hits
	for _, word := range words {
		found = false
		for s, search := range searchlist {
			index, shift := findwordinslice(word, search)
			if index != -1 {
				coords[word] = search[index].coord // Temp
				first := search[index].coord[0]
				second := search[index].coord[1]

				// Add to coords
				switch s {
				case 0: // Top 2 bottom

					first[1] = first[1] + shift            // r + shift
					second[1] = first[1] + (len(word) - 1) // plus length
					coords[word] = [2][2]int{first, second}

				case 1: // Bottom 2 top

					first[1] = first[1] - shift            // r - shift
					second[1] = first[1] - (len(word) - 1) // minus length
					coords[word] = [2][2]int{first, second}

				case 2: // Left 2 right

					first[0] = first[0] + shift            // c + shift
					second[0] = first[0] + (len(word) - 1) // plus length
					coords[word] = [2][2]int{first, second}

				case 3: // Right 2 left

					first[0] = first[0] - shift            // c - shift
					second[0] = first[0] - (len(word) - 1) // minus length
					coords[word] = [2][2]int{first, second}

				case 4: // Top Left 2 Bottom Right

					first[0] = first[0] + shift            // c + shift
					first[1] = first[1] + shift            // r + shift
					second[0] = first[0] + (len(word) - 1) // plus length
					second[1] = first[1] + (len(word) - 1) // plus length

					coords[word] = [2][2]int{first, second}

				case 5: // Bottom Right 2 Top Left

					first[0] = first[0] - shift            // c - shift
					first[1] = first[1] - shift            // r - shift
					second[0] = first[0] - (len(word) - 1) // minus length
					second[1] = first[1] - (len(word) - 1) // minus length

					coords[word] = [2][2]int{first, second}

				case 6: // Top Right 2 Bottom Left

					first[0] = first[0] - shift            // c - shift
					first[1] = first[1] + shift            // r + shift
					second[0] = first[0] - (len(word) - 1) // minus length
					second[1] = first[1] + (len(word) - 1) // plus length

					coords[word] = [2][2]int{first, second}

				case 7: // Bottom Left 2 Top Right

					first[0] = first[0] + shift            // c + shift
					first[1] = first[1] - shift            // r - shift
					second[0] = first[0] + (len(word) - 1) // plus length
					second[1] = first[1] - (len(word) - 1) // minus length

					coords[word] = [2][2]int{first, second}

				default:
					// Nothing
				}

				// Set flag
				found = true
			}
		}
		if !found {
			return nil, errors.New("could not find word: " + word)
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
