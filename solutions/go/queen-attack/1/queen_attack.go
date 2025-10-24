package queenattack

import (
	"math"
	"errors"
)

func diagonal(w1, w2, b1, b2 int) bool {
	return math.Abs(float64(w1) - float64(b1)) == math.Abs(float64(w2) - float64(b2))
}

func CanQueenAttack(whitePosition, blackPosition string) (bool, error) {
	if len(whitePosition) > 2 || len(blackPosition) > 2 || whitePosition == "" || blackPosition == "" {
		return false, errors.New("invalid position")
	}

	w1, w2 := int(whitePosition[0]), int(whitePosition[1])
	b1, b2 := int(blackPosition[0]), int(blackPosition[1])

	// check that queens are on the bounds of the board
	if (w1 < 97 || w1 > 104 || b1 < 97 || b1 > 104 || w2 < 49 || w2 > 56 || b1 < 49 || b2 > 56) {
		return false, errors.New("off board")
	}

	// check that queens are not on the same square
	if whitePosition == blackPosition {
		return false, errors.New("same square")
	}

	if (w1 == b1 || w2 == b2 || diagonal(w1, w2, b1, b2)) {
		return true, nil
	}
	return false, nil
}
