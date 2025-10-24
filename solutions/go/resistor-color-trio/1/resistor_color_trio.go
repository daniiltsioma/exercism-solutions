package resistorcolortrio

import (
	"fmt"
	"math"
)

// intPow calculates powers of integers.
func intPow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}

// Label describes the resistance value given the colors of a resistor.
// The label is a string with a resistance value with an unit appended
// (e.g. "33 ohms", "470 kiloohms").
func Label(colors []string) string {
	colorsToNumbers := map[string]int{
		"black": 	0,
		"brown": 	1,
		"red": 		2,
		"orange": 	3,
		"yellow": 	4,
		"green":	5,
		"blue": 	6,
		"violet": 	7,
		"grey": 	8,
		"white": 	9,
	}

	resistance := 10 * colorsToNumbers[colors[0]]
	resistance += colorsToNumbers[colors[1]]
	resistance *= intPow(10, colorsToNumbers[colors[2]])

	prefix := ""	
	if (resistance > int(1e9)) {
		prefix = "giga"
		resistance /= int(1e9)
	} else if (resistance > int(1e6)) {
		prefix = "mega"
		resistance /= int(1e6)
	} else if (resistance > 1000) {
		prefix = "kilo"
		resistance /= 1000
	}

	return fmt.Sprintf("%d %sohms", resistance, prefix)
}
