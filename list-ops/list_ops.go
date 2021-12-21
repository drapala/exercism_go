package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

// Fold: 
// https://en.wikipedia.org/wiki/Fold_%28higher-order_function%29
// https://burgaud.com/foldl-foldr-python

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	return s[1:].Foldl(fn, fn(initial, s[0]))
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	if len(s) == 0 {
		return initial
	}
	return s[:len(s)-1].Foldr(fn, fn(s[len(s)-1], initial))
}

func (s IntList) Filter(fn func(int) bool) IntList {
	result := make(IntList, 0)
	for _, v := range s {
		if fn(v) {
			result = append(result, v)
		}
	}
	return result
}

func (s IntList) Length() int {
	var count int = 0
	for range s {
		count++
	}
	return count
}

func (s IntList) Map(fn func(int) int) IntList {
	result := make(IntList, 0)
	for _, v := range s {
		result = append(result, fn(v))
	}
	return result
}

func (s IntList) Reverse() IntList {
	result := make(IntList, 0)
	for i := len(s) - 1; i >= 0; i-- {
		result = append(result, s[i])
	}
	return result
}

func (s IntList) Append(lst IntList) IntList {
	var capacity = s.Length() + lst.Length()
	var result = make(IntList, capacity)

	var index int = 0
	// First list
	for _, v := range(s) {
		result[index] = v
		index++
	}
	// Second list
	for _, v := range(lst) {
		result[index] = v
		index++
	}
	return result
}

func (s IntList) Concat(lists []IntList) IntList {
	result := s
	// Iterate over the list of lists
	for _, lst := range(lists) {
		result = result.Append(lst)
	}
	return result
}
