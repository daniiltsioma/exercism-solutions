package pythagorean

type Triplet [3]int

// Range generates list of all Pythagorean triplets with side lengths
// in the provided range.
func Range(min, max int) []Triplet {
	triplets := []Triplet{}

	for a := min; a <= max-2; a++ {
		for b := a+1; b <= max-1; b++ {
			for c := b+1; c <= max; c++ {
				if a*a + b*b == c*c {
					triplets = append(triplets, Triplet{a,b,c})
				}
			}
		}
	}

	return triplets
}

// Sum returns a list of all Pythagorean triplets with a certain perimeter.
func Sum(p int) []Triplet {
	triplets := []Triplet{}
	
	for a := 1; a <= p; a++ {
		for b := a; b <= p; b++ {
			for c := b; c <= p; c++ {
				if (a*a + b*b == c*c) && (a + b + c == p) {
					triplets = append(triplets, Triplet{a,b,c})
				}
			}
		}
	}

	return triplets
}
