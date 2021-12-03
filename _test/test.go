package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	s2 := make([]int, len(s1))

	copy(s2, s1)
	fmt.Println(s1, s2) // [1 2 3 4 5] [1 2 3 4 5]

	s2[1] = 10          // changing s2 does not affect s1
	fmt.Println(s1, s2) // [1 2 3 4 5] [1 10 3 4 5]
}