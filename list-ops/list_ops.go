package listops

import "fmt"

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Fold: 
// https://en.wikipedia.org/wiki/Fold_%28higher-order_function%29
// https://burgaud.com/foldl-foldr-python

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}

	fmt.Println(s)
	// 1: [1 2 3 4]
	// 2: [2 3 4]
	// 3: [3 4]
	// 4: [4]
	// 5: []

	val := s[1:].Foldl(fn, fn(initial, s[0]))
	fmt.Println("Returned:", val)
	return val
	// 1: [2 3 4].Foldl(fn, fn(1, 5))
	// 2: [3 4].Foldl(fn, fn(2, 6))
	// 3: [4].Foldl(fn, fn(3, 8))
	// 4: [].Foldl(fn, fn(4, 11))
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	panic("Please implement the Foldr function")
}

func (s IntList) Filter(fn func(int) bool) IntList {
	panic("Please implement the Filter function")
}

func (s IntList) Length() int {
	panic("Please implement the Length function")
}

func (s IntList) Map(fn func(int) int) IntList {
	panic("Please implement the Map function")
}

func (s IntList) Reverse() IntList {
	panic("Please implement the Reverse function")
}

func (s IntList) Append(lst IntList) IntList {
	panic("Please implement the Append function")
}

func (s IntList) Concat(lists []IntList) IntList {
	panic("Please implement the Concat function")
}
