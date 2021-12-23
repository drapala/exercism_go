package prime

func getDivisors(n int64, divisors *[]int64, testDivisor int64) {
	// If we're at 1, we're done
	if n == 1 {
		return
	}
	if n%testDivisor == 0 { // Divisor is good, try again
		*divisors = append(*divisors, testDivisor)
		getDivisors(n/testDivisor, divisors, testDivisor)
	} else { // Divisor is not good, increment
		// Recurse
		getDivisors(n, divisors, testDivisor+1)
	}
}

func Factors(n int64) []int64 {
	divisors := make([]int64, 0)
	getDivisors(n, &divisors, 2) // Start with 2
	return divisors
}
