// Package weather returns the current weather condition.
package weather

// CurrentCondition returns the current weather condition.
var CurrentCondition string
// CurrentLocation is where you're at.
var CurrentLocation string

// Forecast concatenates stuff.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
