// Package triangle helps determine the characteristics of triangle.
package triangle

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = 0 // Not a triangle
	Equ = 1 // equilateral
	Iso = 2 // isosceles
	Sca = 3 // scalene
)

// KindFromSides determines if a shape is a triangle
// and what type it is based on three sides.
func KindFromSides(a, b, c float64) Kind {
	if a + b <= c  || a + c <= b || b + c <= a {
		return 0
	}

	if a == b && b == c && a == c {
		return 1
	}

	if a == b || b == c || a == c {
		return 2
	}

	return 3
}
