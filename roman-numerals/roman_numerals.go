package romannumerals

import "fmt"

func convert_thousands(digit int) string {
	switch digit {
	case 1:
		return "M"
	case 2:
		return "MM"
	case 3:
		return "MMM"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}

func convert_hundreds(digit int) string {
	switch digit {
	case 1:
		return "C"
	case 2:
		return "CC"
	case 3:
		return "CCC"
	case 4:
		return "CD"
	case 5:
		return "D"
	case 6:
		return "DC"
	case 7:
		return "DCC"
	case 8:
		return "DCCC"
	case 9:
		return "CM"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}

func convert_tens(digit int) string {
	switch digit {
	case 1:
		return "X"
	case 2:
		return "XX"
	case 3:
		return "XXX"
	case 4:
		return "XL"
	case 5:
		return "L"
	case 6:
		return "LX"
	case 7:
		return "LXX"
	case 8:
		return "LXXX"
	case 9:
		return "XC"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}

func convert_ones(digit int) string {
	switch digit {
	case 1:
		return "I"
	case 2:
		return "II"
	case 3:
		return "III"
	case 4:
		return "IV"
	case 5:
		return "V"
	case 6:
		return "VI"
	case 7:
		return "VII"
	case 8:
		return "VIII"
	case 9:
		return "IX"
	case 0:
		return ""
	default:
		return "Invalid digit"
	}
}

func ToRomanNumeral(input int) (string, error) {
	// Error cases
	if input < 1 || input > 3000 {
		return "", fmt.Errorf("number must be between 1 and 3000")
	}

	// Convert number into it's corresponding digits
	ones := input % 10
	tens := input / 10 % 10
	hundreds := input / 100 % 10
	thousands := input / 1000 % 10

	final := convert_thousands(thousands) + convert_hundreds(hundreds) + convert_tens(tens) + convert_ones(ones)

	return final, nil
}
