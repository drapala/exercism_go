package main

import "fmt"

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

func ResultAppend(input string, count int) string {
	// Add to result
	var result string 

	if fmt.Sprint(count) != "1" {
		result += fmt.Sprint(count)
	}
	result += string(input)
	return result
}

func main() {
	input := "AABBBCCCC"
	fmt.Println(input)

	var result string
	var local_index, global_index int

	for i := 0; i < len(input); i++ {
		fmt.Println("Letter of interest: ", string(input[i]), "at i: ", i)
		local_index = findSeqEnd(input[i:], string(input[i]))
		global_index = i + local_index
		fmt.Println("Last ", string(input[i]),  " at global index: ", global_index)
		fmt.Println("Count: ", local_index + 1)
		
		// Add to result
		result += ResultAppend(string(input[i]), local_index + 1)

		i = global_index
	}

	fmt.Println(result)

}
