package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	var result Ints
	for _, value := range i {
		if filter(value) {
			result = append(result, value)
		}
	}
	return result
}

func (i Ints) Discard(filter func(int) bool) Ints {
	var result Ints
	for _, value := range i {
		if !filter(value) {
			result = append(result, value)
		}
	}
	return result
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	var result Lists
	for _, list := range l {
		if filter(list) {
			result = append(result, list)
		}
	}
	return result
}

func (s Strings) Keep(filter func(string) bool) Strings {
	var result Strings
	for _, value := range s {
		if filter(value) {
			result = append(result, value)
		}
	}
	return result
}
