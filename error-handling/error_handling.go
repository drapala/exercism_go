package erratum

// IsTransient returns a boolean indicating whether the error is of type TransientError. It is satisfied by TransientError.
func IsTransient(err error) bool {
	// If no error, return out
	if err == nil {
		return false
	}
	/// Based on type of error, return boolean
	switch err.(type){
	case TransientError:
		return true
	default:
		return false
	}
}

func Use(opener ResourceOpener, input string) (err error) {
	// Open a resource via opener
	resource, errResource := opener()
	if errResource != nil {
		// Keep trying if TransientError
		for IsTransient(errResource) {
			resource, errResource = opener()
			// If not Transient error or nil, immediately return it
			if !IsTransient(errResource) && errResource != nil {
				return (errResource)
			}
		}
	}
	// Defer closing
	defer resource.Close()

	// Defer in case of panic
	defer func() {
		switch p := recover(); p.(type) {
		case nil:
			// no panic
		case FrobError:
			// Defrob using defrobTag
			resource.Defrob(p.(FrobError).defrobTag)
			// Set error for return out of Use
			err = p.(FrobError).inner
		default:
			// Not a FrobError - so return the error
			err = p.(error)
		}
	}()
	// Call Frob(input)
	resource.Frob(input)
	return err
}
