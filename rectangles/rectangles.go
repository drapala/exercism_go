package rectangles

import (
	"fmt"
	"sort"
)

type board struct {
	width, height int
	points        map[coord]string
}

type rectangle struct {
	x [2]int
	y [2]int
}

// ###################################################
// Implement sort on this custom type so that we can sort the co-ordinates
// For implementing Sort - we need interfaces
// See Page 187 in book PDF
type coord struct {
	x, y int
}

type coordslice []coord

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

// ###################################################

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
	//Remove dups from slice.
	inter = removeDups(inter)
	return
}

//Remove duplicates from slice.
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

// 1. Map the co-ordinates of the board in a map
//		Board:
//			Length, Width
//			Co-ordinates
// 2. Loop through all the "+", and for each one, see if a 4-sided rectangle can be formed.
//    Add it to a stack for checking
//    Don't add other rectangles if already exists
// 3. Loop through the stack and check if the rectangle is valid as per the "-" and "|"

func Count(diagram []string) int {
	// Calculate height and width of the board
	height := len(diagram)
	var width int
	points := make(map[coord]string)

	// Print rectangle and form the map of co-ordinates
	for y, line := range diagram {
		if len(line) > width {
			width = len(line)
		}
		for x, char := range line {
			points[coord{x, y}] = string(char)
		}
		fmt.Println(line)
	}

	// 1. Map the co-ordinates of the board in a map
	myBoard := board{
		width:  width - 1,
		height: height - 1,
		points: points,
	}

	fmt.Println("Board:", myBoard)

	// fmt.Println("")

	// fmt.Println("+", myBoard.filter("+"))
	// fmt.Println("-", myBoard.filter("-"))
	// fmt.Println("|", myBoard.filter("|"))

	// 2. Loop through all the "+", and for each one, see if a 4-sided rectangle can be formed.
	// 2a. Do a pass through the list of points and create a hash map, grouping them by x coordinate, so each x coordinate corresponds to a list of y coordinates.
	myMap := make(map[int][]int)
	for _, point := range myBoard.filter("+") {
		myMap[point.x] = append(myMap[point.x], point.y)
	}
	fmt.Println(" ")
	fmt.Println(myMap)

	rectangles := make([]rectangle, 0)

	for x1 := 0; x1 < width; x1++ {
		for x2 := x1 + 1; x2 < width; x2++ {
			if myMap[x1] != nil && myMap[x2] != nil {
				x := [2]int{x1, x2}
				intersections := intersection(myMap[x1], myMap[x2])
				if len(intersections) >= 2 {
					for left := 0; left < len(intersections); left++ {
						for right := left + 1; right < len(intersections); right++ {
							myRectangle := rectangle{
								x: x,
								y: [2]int{intersections[left], intersections[right]},
							}
							rectangles = append(rectangles, myRectangle)
						}
					}
				}
			}
		}
	}

	fmt.Println("valid rectangles: ", rectangles)

	// var validRectanges int

	// // Go across the path to make sure that the rectangle is valid
	// for _, r := range rectangles {
	// 	// X1 --> X2, Y = Y1
	// 	traverseX(r.x[0], r.x[1], r.y[0], points)

	// 	// X1 --> X2, Y = Y2
	// 	traverseX(r.x[0], r.x[1], r.y[1], points)

	// 	// X = X1, Y1 --> Y2
	// 	traverseY(r.x[0], r.y[0], r.y[1], points)

	// 	// X = X2, Y1 --> Y2
	// 	traverseY(r.x[0], r.y[0], r.y[1], points)

	// }

	return len(rectangles)
}

func traverseX(x1, x2, y int, points map[coord]string) bool {
	if x2-x1 == 1 { // X is next to each other
		return true
	}
	// X is not next to each other - ensure everything is a "-"
	for x := x1 + 1; x < x2; x++ {
		if points[coord{x, y}] != "-" {
			return false
		}
	}
	return true
}

func traverseY(x, y1, y2 int, points map[coord]string) bool {
	if y2-y1 == 1 { // Y is next to each other
		return true
	}
	// Y is not next to each other - ensure everything is a "|"
	for y := y1 + 1; y < y2; y++ {
		if points[coord{x, y}] != "|" {
			return false
		}
	}
	return true
}
