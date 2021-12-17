package encode

import (
	"fmt"
	"strconv"
)


func findSeqEnd(text string, letterOfInterest string) int{
	for j := 0; j < len(text); j++ {
		// If we went over to a new letter, return previous index
		if string(text[j]) != letterOfInterest {
			return j-1
		} else {
			// Still have a match
			if j == len(text)-1 {
				// At the end
				return j
			}
		}
	}
	return -1
}

func RLEResultAppend(input string, count int) string {
	// Add to result
	var result string 

	if fmt.Sprint(count) != "1" {
		result += fmt.Sprint(count)
	}
	result += string(input)
	return result
}

func findNumEnd(text string) (int, int) {
	var number int
	for j := 0; j < len(text); j++ {
		// If we went over to a letter, number must have ended at previous index
		if ('0' <= text[j] && text[j] <= '9') {
			continue
		} else {
			// Went over to a letter
			number, _ = strconv.Atoi(text[:j])
			return number, j
		}
	}
	return -1, -1
}

func RLDResultAppend(count int, letter string) string {
	//fmt.Println("Number: ", count, " | String: ", letter )
	var result string
	// Add count number of times
	for i := 1; i <= count; i++{
		result += letter
	}
	return result
}

func RunLengthEncode(input string) string {
	var result string
	var local_index, global_index int

	for i := 0; i < len(input); i++ {
		local_index = findSeqEnd(input[i:], string(input[i]))
		global_index = i + local_index
		// Add to result
		result += RLEResultAppend(string(input[i]), local_index + 1)
		i = global_index
	}
	return result
}

func RunLengthDecode(input string) string {
	var result string
	var count, local_index int

	for i := 0; i < len(input); i++ {
		if ('0' <= input[i] && input[i] <= '9') {
			// Find end of number
			count, local_index = findNumEnd(input[i:])
			// Increment i to start after local_index
			i += local_index
			// Next must be a single letter
			result += RLDResultAppend(count, string(input[i]))
        } else {
			// Got the string directly
			result += RLDResultAppend(1, string(input[i]))
		}
	}
	return result
}
