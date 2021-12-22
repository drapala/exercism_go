package armstrong

import (
	"math"
	"strconv"
)

func numdigits(n int) int {
	s := strconv.Itoa(n)
	return len(s)
}

func calcArmstrongValue(n, length int) int {
	s := strconv.Itoa(n)
	var result float64
	for _, c := range s {
		v, _ := strconv.Atoi(string(c))
		result += math.Pow(float64(v), float64(length))
	}
	return int(result)
}

func IsNumber(n int) bool {
	return calcArmstrongValue(n, numdigits(n)) == n
}
