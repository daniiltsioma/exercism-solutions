package diamond

import (
	"strings"
	"errors"
)

// Gen generates a diamond for specified char.
func Gen(char byte) (string, error) {
	if char < 'A' || char > 'Z'	{
		return "", errors.New("Invalid character")
	}

	if char == byte('A') {
		return "A", nil
	}
	
	output := ""

	width := 2 * int(char - 'A') + 1
	sideSpaces := (width - 1) / 2
	
	// top row
	output += strings.Repeat(" ", sideSpaces)
	output += "A"
	output += strings.Repeat(" ", sideSpaces)
	output += "\n"
	
	sideSpaces -= 1
	midSpaces := 1

	// continue top half
	for c := byte('B'); c < char; c++ {
		output += strings.Repeat(" ", sideSpaces)
		output += string(c)
		output += strings.Repeat(" ", midSpaces)
		output += string(c)
		output += strings.Repeat(" ", sideSpaces)
		output += "\n"

		sideSpaces -= 1
		midSpaces += 2
	}

	// mid row
	output += string(char)
	output += strings.Repeat(" ", midSpaces)
	output += string(char)
	output += "\n"

	sideSpaces += 1
	midSpaces -= 2

	// continue bottom half
	for c := char - 1; byte('A') < c; c-- {
		output += strings.Repeat(" ", sideSpaces)
		output += string(c)
		output += strings.Repeat(" ", midSpaces)
		output += string(c)
		output += strings.Repeat(" ", sideSpaces)
		output += "\n"

		sideSpaces += 1
		midSpaces -= 2
	}

	// bottom row
	output += strings.Repeat(" ", sideSpaces)
	output += "A"
	output += strings.Repeat(" ", sideSpaces)

	return output, nil
}