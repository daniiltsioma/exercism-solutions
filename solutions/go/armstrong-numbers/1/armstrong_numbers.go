package armstrong

import (
	"strconv"
	"math"
)

func IsNumber(n int) bool {
	digits := strconv.Itoa(n)
	p := len(digits)

	sum := 0
	for _, d := range digits {
		sum += int(math.Pow(float64(int(d - '0')), float64(p)))
	}	

	return sum == n
}
