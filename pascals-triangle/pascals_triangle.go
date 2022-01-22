package pascal

func Triangle(n int) [][]int {
	result := make([][]int, n)
	// Loop from 1 to n
	for i := 1; i <= n; i++ {
		row := make([]int, i)    // i is the length of this new row
		for j := 0; j < i; j++ { // Loop over current row
			if j == 0 || j == i-1 { // Leftmost, rightmost is 1
				row[j] = 1
			} else { // Somewhere in the middle
				prevRow := result[i-2]   // i-2 because zero index, and we go back one more
				row[j] = prevRow[j-1] + prevRow[j]
			}
		}
		result[i-1] = row
	}
	return result
}
