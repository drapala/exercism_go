package space

type Planet string

func Age(seconds float64, planet Planet) float64 {
	// An Earth year in seconds
	var EarthYear float64 = 31557600

	// Multiplier based on planet
	var multiplier float64

	switch planet {
	case "Mercury":
		multiplier = 0.2408467
	case "Venus":
		multiplier = 0.61519726
	case "Earth":
		multiplier = 1
	case "Mars":
		multiplier = 1.8808158
	case "Jupiter":
		multiplier = 11.862615
	case "Saturn":
		multiplier = 29.447498
	case "Uranus":
		multiplier = 84.016846
	case "Neptune":
		multiplier = 164.79132
	}
	var PlanetAge float64
	PlanetAge = seconds / (EarthYear * multiplier)

	return PlanetAge
}
