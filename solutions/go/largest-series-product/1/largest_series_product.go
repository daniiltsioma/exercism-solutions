package lsproduct

import (
	"math"
	"errors"
	"unicode"
)

// rti converts a rune into a int64 number.
func rti(digit rune) int64 {
	return int64(digit - '0')
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("span must be smaller than string length")
	}
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}

	currentMax := int64(0)
	ephProd := int64(0)
	zeros := 0
	
	for i := 0; i < span; i++ {
		// check if rune is a digit
		r := rune(digits[i])
		if !unicode.IsDigit(r) {
			return 0, errors.New("digits input must only contain digits")
		}

		d := rti(r)
		if d == 0 {
			zeros++
		} else if ephProd == 0 {
			ephProd = d
		} else {
			ephProd *= d
		}
		if zeros <= 0 {
			currentMax = int64(math.Max(float64(currentMax), float64(ephProd)))
		} else {
			currentMax = 0
		}
	}

	for i := span; i < len(digits); i++ {
		// check if rune is a digit
		r := rune(digits[i])
		if !unicode.IsDigit(r) {
			return 0, errors.New("digits input must only contain digits")
		}


		l := rti(rune(digits[i-span]))
		d := rti(rune(digits[i]))

		if l == 0 {
			zeros--
		} else {
			ephProd /= l
		}
		
		if d == 0 {
			zeros++
		} else if ephProd == 0 {
			ephProd = d
		} else {
			ephProd *= d
		}
		if zeros <= 0 {
			currentMax = int64(math.Max(float64(currentMax), float64(ephProd)))
		} 
	}	

	return currentMax, nil
}
