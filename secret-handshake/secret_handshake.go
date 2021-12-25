package secret

import "fmt"

// Calculate binary representation of integer
func binary(n uint) string {
	var result string
	for n > 0 {
		result = fmt.Sprintf("%d%s", n%2, result)
		n /= 2
	}
	return result
}
// Generate secret list from binary code
func generateList(binaryCode string) []string {
	result := make([]string, 0)
	for i:=len(binaryCode)-1; i>=0; i-- {
		if binaryCode[i] == '1' {
			switch i {
			case len(binaryCode)-1:
				result = append(result, "wink")
			case len(binaryCode)-2:
				result = append(result, "double blink")
			case len(binaryCode)-3:
				result = append(result, "close your eyes")
			case len(binaryCode)-4:
				result = append(result, "jump")
			}
		}
	}
	return result
}
// Reverse array and export
func reverseArray(input[]string) []string{
	var output []string
	for i:=len(input)-1; i>=0; i-- {
		output = append(output, input[i])
	}
	return output
}

func Handshake(code uint) []string {
	result := make([]string, 0)
	// Calculate binary representation of integer
	binaryCode := binary(code)
	var reverse bool
	// Check if we need to reverse order
	if len(binaryCode) == 5 {
		reverse = true
		binaryCode = binaryCode[1:] // Trim to actual message bits
	}
	// Generate secret list from binary code
	result = generateList(binaryCode)
	// Reverse array if needed
	if reverse {
		result = reverseArray(result)
	}
	return result
}
