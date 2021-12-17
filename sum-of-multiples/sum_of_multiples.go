package summultiples

func SumMultiples(limit int, divisors ...int) int {
	var multiples []int
	for n := 1; n < limit; n++ {
		// Attempt to divide n by each divisor
		for _, divisor := range divisors {
			if divisor != 0{
				if n%divisor == 0 {
					multiples = append(multiples, n)
					break
				}
			}
		}
	}
	var sum int
	for _, multiple := range multiples {
		sum += multiple
	}
	return sum
}
