package prime

func IsPrime(x int) bool {
	// If the number is divisble only by itself and 1, it is prime
	if x == 2 {
		return true
	}
	for i := 2; i < x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func Nth(n int) (int, bool) {
	// Error handling
	if n < 1 {
		return 0, false
	}

	var num int = 0
	for x := 2; x <= 2147483647; x++ {
		if IsPrime(x) {
			num++
			if num == n {
				return x, true
			}
		}
	}
	return 0, false
}
