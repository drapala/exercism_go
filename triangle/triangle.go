package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind string

const (
	// Pick values for the following identifiers used by the test program.
	NaT Kind = "NaT" // not a triangle
	Equ Kind = "Equ" // equilateral
	Iso Kind = "Iso" // isosceles
	Sca Kind = "Sca" // scalene
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind

	// Check if Triangle
	if a <= 0 || b <= 0 || c <= 0 {
		// Negative or zero sides
		k = NaT
		return k
	} else if (a+b < c) || (a+c < b) || (b+c < a) {
		// Triangle inequality fails
		k = NaT
		return k
	}

	// Triangle logic
	if a == b && b == c {
		// Equilateral
		k = Equ
	} else if (a == b && b != c) || (a == c && b != c) || (b == c && a != c) {
		// Isosceles
		k = Iso
	} else {
		// Scalene
		k = Sca
	}

	return k
}
