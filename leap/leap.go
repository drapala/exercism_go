// Calculates if a given year is a leap year
package leap

// Returns true if year is a leap year, otherwise returns false
func IsLeapYear(year int) bool {
	// If year is divisible by 4, it is a leap year
	if year%4 == 0 {
		// But if year is divisible by 100, it might not be a leap year
		if year%100 == 0 {
			// Unless that year is divisible by 400, then it is a leap year
			if year%400 == 0 {
				// So it's a leap year
				return true
			}
			// Not a leap year
			return false
		}
		// It's a leap year
		return true
	}
	// It's not even divisible by 4, so it's not a leap year
	return false
}
