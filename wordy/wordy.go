package wordy

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Errors
var (
    ErrOpNotSupported  = fmt.Errorf("operation not supported")
)

func TrimCommand(input string) string {
	inter := regexp.MustCompile(`What is `).ReplaceAllString(input, "")
	output := regexp.MustCompile(`\?`).ReplaceAllString(inter, "")
	output = strings.ReplaceAll(output, " ", "")
	return output
}

func getDigitsAndCommands(input string) ([]string, []string) {
	digits := make([]string, 0)
	commands := make([]string, 0)

	matches := regexp.MustCompile("-[0-9]+|[0-9]+").FindAllStringIndex(input, -1)

	// Fill up digits
	for i := 0; i < len(matches); i++ {
		digits = append(digits, input[matches[i][0]:matches[i][1]])
	}

	// Fill up commands
	for i := 0; i < len(matches) - 1; i++ {
		commands = append(commands, input[matches[i][1]:matches[i+1][0]])
	}

	return digits, commands
}

func calculator(digits, commands []string) (int, error) {
	result := 0
	for i := 0; i < len(commands); i++ {
		var firstDigit, secondDigit int
		if i == 0 {
			firstDigit, _ = strconv.Atoi(digits[i]) // First op
		} else {
			firstDigit = result // Subsequent op to use result
		}
		secondDigit, _ = strconv.Atoi(digits[i+1])

		switch commands[i] {
		case "plus":
			result = firstDigit + secondDigit
		case "minus":
			result = firstDigit - secondDigit
		case "multipliedby":
			result = firstDigit * secondDigit
		case "dividedby":
			result = firstDigit / secondDigit
		default: // Operation we don't support
			return 0, ErrOpNotSupported
		}
	}
	return result, nil
}

func justNumber(input string, digits []string) bool {
	if len(digits) == 0 {
		return false // Error case - handle outside
	}
	if input == digits[0] {
		return true // Entire input is just a number
	}
	return false
}

func continuousDigits(input string) bool {
	re := regexp.MustCompile("[0-9]+ +[0-9]+")// 2 numbers seperated by space
	matches := re.FindAllStringIndex(input, -1)

	if len(matches) != 0 {
		return true
	} else {
		return false
	}
}

func Answer(question string) (int, bool) {
	output := TrimCommand(question)
	digits, commands := getDigitsAndCommands(output)
	calculated, err := calculator(digits, commands)

	// Handle just a number
	if justNumber(output, digits) {
		digit, _ := strconv.Atoi(output)
		return digit, true
	}

	// Handle continuous Digits
	if continuousDigits(question) {
		return 0, false
	}

	// Handle error
	if err != nil {
		if err == ErrOpNotSupported {
			return 0, false
		}
	}

	// Now - since we've handled just a number, we can assume that commands can be relied on for reliable error handling
	if len(commands) == 0 {
		return 0, false
	}

	return calculated, true
}
