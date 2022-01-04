package main

import "fmt"

func main() {
	half_index := 7 / 2
	my_list := []int{1, 3, 4, 6, 8, 9, 11}

	left := my_list[:half_index]
	right := my_list[half_index:]

	fmt.Println("left", left)
	fmt.Println("right", right)
}