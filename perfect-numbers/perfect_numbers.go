package perfect

import "errors"

// Define the Classification type here.
type Classification string

const ClassificationAbundant Classification = "ClassificationAbundant"
const ClassificationDeficient Classification = "ClassificationDeficient"
const ClassificationPerfect Classification = "ClassificationPerfect"

// Errors
var ErrOnlyPositive error = errors.New("only positive numbers are allowed")

// Find all positive divisors for a number n
func getDivisors(n int) []int {
	result := make([]int, 0)
	for x := 1; x < n; x++ { // Find divisors
		if n % x == 0 { // Divides cleanly
			result = append(result, x)
		}
	}
	return result
}

// Sums up the array
func getDivisorSum(array []int) int {
	var sum int
	for _, divisor := range(array) {
		sum += divisor
	}
	return sum
}

func Classify(n int64) (Classification, error) {
	// Error handling
	if n <= 0 {
		return "", ErrOnlyPositive
	}

	divisors := getDivisors(int(n))
	sum := getDivisorSum(divisors)

	if sum == int(n) {
		return ClassificationPerfect, nil
	} else if sum > int(n) {
		return ClassificationAbundant, nil
	} else {
		return ClassificationDeficient, nil
	}
}
