package triangle

import "math"

// Kind is representative of the type of triangle
type Kind int

// constants are losely typed until they are used in a strict context
const (
	NaT Kind = iota // not a triangle
	Equ             // equilateral
	Iso             // isosceles
	Sca             // scalene
)

// Takes a triangle and checks if it's either the kind of triangle passed in or NaT
func whatKind(sides []float64, tri Kind) Kind {
	kind := tri
	a := sides[0]
	b := sides[1]
	c := sides[2]

	// Checks to see if the sides are valid numbers
	for _, side := range sides {
		if (side <= 0) || (math.IsNaN(side) == true) || (math.IsInf(side, 0) == true) {
			kind = NaT
			break
		}
	}

	// If any combinataion of 2 sides is less than the 3rd side then return NaT
	if (a+b < c) || (a+c < b) || (b+c < a) {
		kind = NaT
	}

	return kind
}

// KindFromSides uses the side lengths provided to decide what kinds of triangle the sides make up
func KindFromSides(a, b, c float64) Kind {
	var kind Kind

	sides := []float64{a, b, c}

	// If all sides are equal (Equilateral)
	if (a == b) && (b == c) && (a == c) {
		return whatKind(sides, Equ)
	}

	// If two of the sides are equal (Isosceles)
	if (a == b) || (b == c) || (a == c) {
		return whatKind(sides, Iso)
	}

	// If none of the sides are equal (Scalene)
	if (a != b) && (b != c) && (a != c) {
		return whatKind(sides, Sca)
	}

	return kind
}
