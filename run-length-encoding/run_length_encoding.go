package encode

import "fmt"

func ResultAppend(input string, count int) string {
	// Add to result
	var result string 

	if fmt.Sprint(count) != "1" {
		result += fmt.Sprint(count)
	}
	result += string(input)
	return result
}

func RunLengthEncode(input string) string {
	fmt.Println("input: ", input)
	var result string
	var slow, fast, count int
	for slow = 0; slow < len(input); slow++ {
		fmt.Println("slow: ", slow, " | ", string(input[slow]))
		count = 1
		// Loop over and find number of consecutive characters
		for fast = slow + 1; fast < len(input); fast++ {
			fmt.Println("fast: ", fast, " | ", string(input[fast]))
			if input[slow] == input[fast] {
				fmt.Println("Found ", string(input[fast]), " at ", fast)
				count++
				// At the end
				if fast == len(input){
					result += ResultAppend(string(input[fast]), count)
				}
			} else {
				// Add to result
				result += ResultAppend(string(input[fast]), count)
				
				// Set slow to next index as it will increment again
				slow = fast - 1
				fmt.Println("moving slow: ", slow, " | ", string(input[slow]))
				break
			}
		}
	}
	fmt.Println("output:", result)
	return result
}

func RunLengthDecode(input string) string {
	var result string
	return result
}
