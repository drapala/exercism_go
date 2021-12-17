package isbn

import (
	"strconv"
	"strings"
)

func ContainsAllValidChars(isbn string) bool {
	for i := 0; i < len(isbn); i++ {
		if ('0' <= isbn[i] && isbn[i] <= '9') || ('X' == isbn[i]) {
			continue
		} else {
			// Got an invalid character
			return false
		}
	}
	// Went through loop without hitting invalid characters
	return true
}

func ValidISBNLength(isbn string) bool {
	if len(isbn) == 10 {
		return true
	} else {
		return false
	}
}

func ValidISBNChecksum(isbn string) bool {
	x1, _ := strconv.Atoi(string(isbn[0]))
	x2, _ := strconv.Atoi(string(isbn[1]))
	x3, _ := strconv.Atoi(string(isbn[2]))
	x4, _ := strconv.Atoi(string(isbn[3]))
	x5, _ := strconv.Atoi(string(isbn[4]))
	x6, _ := strconv.Atoi(string(isbn[5]))
	x7, _ := strconv.Atoi(string(isbn[6]))
	x8, _ := strconv.Atoi(string(isbn[7]))
	x9, _ := strconv.Atoi(string(isbn[8]))

	var x10 int
	// Special case for x10
	if isbn[9] == 'X' {
		x10 = 10
	} else {
		x10, _ = strconv.Atoi(string(isbn[9]))
	}

	//(x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) mod 11
	checksum := (x1 * 10 + x2 * 9 + x3 * 8 + x4 * 7 + x5 * 6 + x6 * 5 + x7 * 4 + x8 * 3 + x9 * 2 + x10 * 1) % 11

	if checksum == 0 {
		return true
	} else {
		return false
	}
}

func IsValidISBN(isbn string) bool {
	// replace dashes in isbn
	isbn = strings.Replace(isbn, "-", "", -1)
	// Check if isbn contains invalid characters
	if !ContainsAllValidChars(isbn) {
		return false
	}
	// Check if isbn is too long/short
	if !ValidISBNLength(isbn) {
		return false
	}
	// Check if checksum is valid
	if !ValidISBNChecksum(isbn) {
		return false
	}
	// Must be valid ISBN
	return true
}
