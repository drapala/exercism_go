package purchase

import "fmt"

// NeedsLicense determines whether a license is need to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	return kind == "car" || kind == "truck"
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
		return fmt.Sprintf(option1 + " is clearly the better choice.")
	} else {
		return fmt.Sprintf(option2 + " is clearly the better choice.")
	}
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		return originalPrice * 0.80
	} else if age >= 10 {
		return originalPrice * 0.50
	} else if age > 3 && age < 10 {
		return originalPrice * 0.70
	} else {
		return 0.0
	}
}
