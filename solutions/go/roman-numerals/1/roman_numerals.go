package romannumerals

import (
	"strings"
	"errors"
)

func ToRomanNumeral(input int) (string, error) {
	if input < 1 || input > 3999 {
		return "", errors.New("Invalid numeral")
	}

	res := ""

	thousands := input / 1000
	res += strings.Repeat("M", thousands)
	input -= thousands * 1000

	hundreds := input / 100
	if hundreds == 9 {
		res += "CM"
	} else if hundreds >= 5 {
		res += "D" + strings.Repeat("C", hundreds - 5)
	} else if hundreds == 4 {
		res += "CD"
	} else {
		res += strings.Repeat("C", hundreds)
	}
	input -= hundreds * 100

	tens := input / 10
	if tens == 9 {
		res += "XC"
	} else if tens >= 5 {
		res += "L" + strings.Repeat("X", tens - 5)
	} else if tens == 4 {
		res += "XL"
	} else {
		res += strings.Repeat("X", tens)
	}
	input -= tens * 10

	if input == 9 {
		res += "IX"
	} else if input >= 5 {
		res += "V" + strings.Repeat("I", input - 5)
	} else if input == 4 {
		res += "IV"
	} else {
		res += strings.Repeat("I", input)
	}

	return res, nil
}
