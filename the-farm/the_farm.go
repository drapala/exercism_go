package thefarm

import (
	"errors"
	"fmt"
)

// See types.go for the types defined for this exercise.
type SillyNephewError struct {
	cows int
  }
  
func (e *SillyNephewError) Error() string {
	return fmt.Sprintf("silly nephew, there cannot be %d cows", e.cows)
}

// DivideFood computes the fodder amount per cow for the given cows.
func DivideFood(weightFodder WeightFodder, cows int) (float64, error) {
	// So we're receiving the WeightFodder interface, whereas the actual struct testWeightFodder is defined in the_farms_test.go, and so is the method (and that's ok - the interface only cares about implementing a method named FodderAmount() - it doesn't care about where it's defined or which struct it comes from.)
	// Since our tests are coming in as WeightFodder, and testWeightFodder satisfies the method, that's how the relationship is formed it seems
	fodder, err := weightFodder.FodderAmount()
	
	// Handle negative fodder
	if fodder < 0 {
		return 0.0, errors.New("Negative fodder") 
	}

	// Handle division by zero
	if cows == 0 {
		return 0.0, errors.New("Division by zero") 
	}

	// Handle negative cows
	if cows < 0 {
		return 0.0, &SillyNephewError{cows: cows}
	}

	// Error handling
	if err == ErrScaleMalfunction {
		// Negative fodder is handled above already
		return (fodder * 2) / float64(cows), nil
	} else if err != nil && err != ErrScaleMalfunction {
		// For any other error, return 0 and the error.
		return 0, err
	} 

	// Happy Path
	return fodder / float64(cows), nil
}
