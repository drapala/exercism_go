package rectangles

import (
	"sort"
)

// The game board
type board struct {
	width, height int
	points        map[coord]string
}

// Returns a list of co-ordinates for the desired character
func (b *board) filter(c string) coordslice {
	var filtered coordslice
	for k, v := range b.points {
		if v == c {
			filtered = append(filtered, k)
		}
	}
	sort.Sort(filtered)
	return filtered
}

// Cartesian co-ordinates
type coord struct {
	x, y int
}

// Slice of co-ordinates, used for sorting
type coordslice []coord

// Represents a found rectangle
type rectangle struct {
	x [2]int
	y [2]int
}

// Implement sort on this custom type so that we can sort the co-ordinates
func (c coordslice) Len() int { return len(c) }
func (c coordslice) Less(i, j int) bool {
	// Prioritize X first
	if c[i].x < c[j].x {
		return true
	} else if c[i].x > c[j].x {
		return false
	} else { // If X is equal, prioritize Y
		return c[i].y < c[j].y
	}
}
func (c coordslice) Swap(i, j int) { c[i], c[j] = c[j], c[i] }

// Find the intersection of two lists
func intersection(s1, s2 []int) (inter []int) {
	hash := make(map[int]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		// If elements present in the hashmap then append intersection list.
		if hash[e] {
			inter = append(inter, e)
		}
	}
	//Remove duplicates from slice
	inter = removeDups(inter)
	return
}

// Remove duplicates from slice.
func removeDups(elements []int) (nodups []int) {
	encountered := make(map[int]bool)
	for _, element := range elements {
		if !encountered[element] {
			nodups = append(nodups, element)
			encountered[element] = true
		}
	}
	return
}

// Traverses the X axis and checks for valid characters
func traverseX(x1, x2, y int, points map[coord]string) bool {
	// Ensure everything in path is a "-" or "+"
	for x := x1 + 1; x < x2; x++ {
		if points[coord{x, y}] != "-" && points[coord{x, y}] != "+" {
			return false
		}
	}
	return true
}

// Traverses the Y axis and checks for valid characters
func traverseY(x, y1, y2 int, points map[coord]string) bool {
	// Ensure everything in path is a "|" or "+"
	for y := y1 + 1; y < y2; y++ {
		if points[coord{x, y}] != "|" && points[coord{x, y}] != "+" {
			return false
		}
	}
	return true
}

// Counts the number of valid rectangles as per the rules of drawing
func Count(diagram []string) int {
	// Calculate height and width of the board
	height := len(diagram)
	var width int
	points := make(map[coord]string)

	// Form a map of co-ordinates
	for y, line := range diagram {
		if len(line) > width {
			width = len(line)
		}
		for x, char := range line {
			points[coord{x, y}] = string(char)
		}
	}

	// 1. Map the co-ordinates of the board in a board object
	myBoard := board{
		width:  width - 1,
		height: height - 1,
		points: points,
	}

	// 2. Loop through all the "+", and for each one, see if a 4-sided rectangle can be formed.
	groupByX := make(map[int][]int)    // Map to group the co-ordinates by X for easier calculations
	rectangles := make([]rectangle, 0) // To store found rectangles
	// Do a pass through the list of points and create a hash map, grouping them by x coordinate, so each x coordinate corresponds to a list of y coordinates.
	for _, point := range myBoard.filter("+") {
		groupByX[point.x] = append(groupByX[point.x], point.y)
	}
	// Loop through the groups of Y's for possible X's and look for intersections
	// More than 2 intersections means we have rectangle(s)
	for x1 := 0; x1 < width; x1++ {
		for x2 := x1 + 1; x2 < width; x2++ {
			if groupByX[x1] != nil && groupByX[x2] != nil { // Valid X valies
				x := [2]int{x1, x2}
				intersections := intersection(groupByX[x1], groupByX[x2]) // Intersection of the Y's
				if len(intersections) >= 2 {                              // If rectangle
					for left := 0; left < len(intersections); left++ {
						for right := left + 1; right < len(intersections); right++ {
							myRectangle := rectangle{
								x: x,
								y: [2]int{intersections[left], intersections[right]},
							}
							rectangles = append(rectangles, myRectangle) // Append to the list of found rectangles
						}
					}
				}
			}
		}
	}

	// 3. Loop through the stack and check if the rectangle is valid as per the "-" and "|"
	var validRectangles int // The number of valid rectangles from the found rectangles
	for _, r := range rectangles {
		// X1 --> X2, Y = Y1, X1 --> X2, Y = Y2, X = X1, Y1 --> Y2, X = X2, Y1 --> Y2
		if traverseX(r.x[0], r.x[1], r.y[0], points) && traverseX(r.x[0], r.x[1], r.y[1], points) && traverseY(r.x[0], r.y[0], r.y[1], points) && traverseY(r.x[1], r.y[0], r.y[1], points) {
			validRectangles++
		}
	}
	return validRectangles
}
