package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	var result = make(map[string]int)

	// Loop through the keys and values inside in
	for key, value := range in {
		for _, letter := range value {
			result[strings.ToLower(letter)] = key
		}
	}
	return result
}
