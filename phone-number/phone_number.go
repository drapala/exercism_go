package phonenumber

import (
	"errors"
	"fmt"
)

func CleanNumber(phoneNumber string) (string, error) {
	var result string

	for _, c := range(phoneNumber){		
		// If c is a valid numeral
		if c >= '0' && c <= '9' {
			result += string(c)
		}
	}
	return result, nil
}

func CheckValidN(digit rune) bool {
	if digit >= '2' && digit <= '9' {
		return true
	} else {
		return false
	}
}

func Number(phoneNumber string) (string, error) {
	clean, _ := CleanNumber(phoneNumber)

	// Error handling
	if len(clean) < 10 { // Minimum 10 digits: NXX NXX XXXX
		return "", errors.New("invalid - number is less than 10 digits")
	} else if len(clean) > 11 { // Maximum 11 digits: 1 NXX NXX XXXX
		return "", errors.New("invalid - number is greater than 11 digits")
	}

	// Handle 11 digits case to homogenize remainder of operations
	if len(clean) == 11 {
		if clean[0] != '1' {
			return "", errors.New("invalid - country code must be 1")
		}
		clean = clean[1:] // Trim the country code to homogenize
	}
	
	// Handle invalid N: NXX NXX XXXX
	N_area := clean[0]
	N_exchange := clean[3]
	
	if !CheckValidN(rune(N_area)) || !CheckValidN(rune(N_exchange)) {
		return "", errors.New("invalid N - '2' <= N <= '9'")
	}

	return clean, nil
}

func AreaCode(phoneNumber string) (string, error) {
	clean, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return clean[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	clean, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", clean[0:3], clean[3:6], clean[6:]), nil
}
