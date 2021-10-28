package clock

import "fmt"

// Create Clock Struct
type Clock struct {
	hour, minute int
}

// Convert to string
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// New Clock
func New(hours int, minutes int) Clock {
	// Turn everything into minutes
	minutes = hours * 60 + minutes

	// Get hours and mins
	hours = (minutes / 60) % 24
	minutes = minutes % 60

	// Deal with negatives now
	if minutes < 0 {
		hours -= 1 // Roll back one hour
		minutes += 60 // Make mins positive
	}
	if hours < 0 {
		hours += 24 // Normalize to 24 hours
	}

	return Clock{hours, minutes}
}

// Add Minutes
func (c Clock) Add(minutes int) Clock {
	return New(c.hour, c.minute + minutes)
}

// Subtract Minutes
func (c Clock) Subtract(minutes int) Clock {
	return New(c.hour, c.minute - minutes)
}