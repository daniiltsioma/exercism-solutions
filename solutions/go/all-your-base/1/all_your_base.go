package allyourbase

import (
	"math"
	"errors"
)

// intPow calculates powers of integers.
func intPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	// return zero for an empty slice.
	if len(inputDigits) == 0 {
		return []int{0}, nil
	}

	// convert to decimal.
	dec := 0
	for _, d := range inputDigits {
		if d < 0 || d >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
		dec = dec * inputBase + d
	}

	// zero is zero in all bases.
	if dec == 0 {
		return []int{0}, nil
	}

	// convert to output base.
	outputDigits := []int{}
	for dec > 0 {
		rem := dec % outputBase
		outputDigits = append([]int{rem}, outputDigits...)
		dec /= outputBase
	}

	return outputDigits, nil
}
