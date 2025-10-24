package perfect

import (
	"errors"
)

// Define the Classification type here.
type Classification int

const (
	ClassificationDeficient Classification = iota + 1
	ClassificationPerfect
	ClassificationAbundant
)

var ErrOnlyPositive = errors.New("Non-positive numbers not allowed")

func Classify(n int64) (Classification, error) {
	if n < 1 {
		return 0, ErrOnlyPositive
	}

	factors := []int64{}

	for i := int64(1); i < n; i++ {
		if n % i == 0 {
			factors = append(factors, i)
		}
	}

	aliquotSum := int64(0)
	for _, f := range factors {
		aliquotSum += int64(f)
	}

	if aliquotSum < n {
		return ClassificationDeficient, nil
	}
	if n < aliquotSum {
		return ClassificationAbundant, nil
	}

	return ClassificationPerfect, nil
}
