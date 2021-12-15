package main

import "fmt"

type Ints []int
type Lists [][]int
type Strings []string

// Ints
func lt10(x int) bool { return x < 10 }
func gt10(x int) bool { return x > 10 }
func odd(x int) bool  { return x&1 == 1 }
func even(x int) bool { return x&1 == 0 }

// Strings
func zword(s string) bool { return len(s) > 0 && s[0] == 'z' }

// Lists
func has5(l []int) bool {
	for _, e := range l {
		if e == 5 {
			return true
		}
	}
	return false
}

func (i Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, value := range(i){
		if filter(value){
			result = append(result, value)
		}
	}
	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var result Ints
	for _, value := range(i){
		if !filter(value){
			result = append(result, value)
		}
	}
	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, value := range(s){
		if filter(value){
			result = append(result, value)
		}
	}
	return result
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, list := range(l){
		if filter(list){
			result = append(result, list)
		}
	}
	return result
}

func main() {
	// Ints
	nums := Ints{1, 2, 3, 9, 10, 11}
	fmt.Println(nums.Keep(lt10))
	fmt.Println(nums.Discard(lt10))
	fmt.Println(nums.Keep(odd))
	fmt.Println(nums.Keep(even))

	// Strings
	string_list := Strings{"apple", "zebra", "banana", "zombies", "cherimoya", "zealot"}
	fmt.Println(string_list.Keep(zword))

	// Lists
	list := Lists{
		{1, 2, 3},
		{5, 5, 5},
		{5, 1, 2},
		{2, 1, 2},
		{1, 5, 2},
		{2, 2, 1},
		{1, 2, 5},
	}
	fmt.Println(list.Keep(has5))
}
