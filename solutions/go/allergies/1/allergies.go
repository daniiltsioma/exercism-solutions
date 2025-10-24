package allergies

import (
	"math"
)

var knownAllergens = []string{
	"eggs",
	"peanuts",
	"shellfish",
	"strawberries",
	"tomatoes",
	"chocolate",
	"pollen",
	"cats",
}

// intPow calculates powers of unsigned integers
func uintPow(a, b uint) uint {
	return uint(math.Pow(float64(a), float64(b)))
}

func Allergies(allergies uint) []string {
	// handle the case of no allergies.
	if allergies == 0 {
		return []string{}
	}

	// normalize allergies input.
	if allergies >= 256 {
		allergies %= 256
	}

	res := []string{}

	for i := len(knownAllergens) - 1; i >= 0; i-- {
		allergyScore := uintPow(2, uint(i))
		if allergies / allergyScore == 1 {
			res = append([]string{knownAllergens[i]}, res...)
			allergies %= allergyScore
		}
	}

	return res
}

func AllergicTo(allergies uint, allergen string) bool {
	detectedAllergies := Allergies(allergies)
	for _, a := range detectedAllergies {
		if a == allergen {
			return true
		}
	}
	return false
}
