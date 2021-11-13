package elon

import "fmt"

// Drive the car
func (car *Car) Drive() {
	// If car has enough battery for one more drive
	if car.battery >= car.batteryDrain {
		car.battery -= car.batteryDrain
		car.distance += car.speed
	}
}

// Display distance
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// Display battery
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// TODO: define the 'CanFinish(trackDistance int) bool' method
func (car *Car) CanFinish(trackDistance int) bool {
	num_drives := car.battery / car.batteryDrain
	actualDistance := car.speed * num_drives

	if actualDistance >= trackDistance {
		return true
	} else {
		return false
	}
}