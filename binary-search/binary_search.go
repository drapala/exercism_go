package binarysearch

import "fmt"

func SearchInts(list []int, key int) int {
	fmt.Println("list:", list, "key:", key)

	if list == nil { // Called on empty list - key not present
		return -1
	}
	var half_index int = len(list) / 2
	if key > list[half_index] {
		return SearchInts(list[half_index:], key)
	} else if key < list[half_index] {
		return SearchInts(list[:half_index], key)
	}
	return half_index
}
