package collatzconjecture

import "fmt"

func CollatzConjecture(n int) (int, error) {
	var count int

	if n <= 0 {
		return 0, fmt.Errorf("%d is not a positive number", n)
	}

	// Loop until we get to 1
	for n != 1 {
		count++
		if n%2 == 0 { // Even
			n = n / 2
		} else { // Odd
			n = 3*n + 1
		}
	}
	return count, nil
}
