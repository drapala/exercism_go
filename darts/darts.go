package darts

import "math"

func Score(x, y float64) int {
	// Find radius of point
	var r float64 = math.Sqrt(x*x + y*y)

	switch {
	case r <= 1.0:
		// Inner circle
		return 10
	case r <= 5.0 && r > 1.0:
		// Middle circle
		return 5
	case r <= 10.0 && r > 5.0:
		// Outer circle
		return 1
	default:
		// Missed target
		return 0
	}
}
