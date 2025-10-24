package encode

import (
	"unicode"
	"strings"
	"strconv"
)

func RunLengthEncode(input string) string {
	if input == "" {
		return ""
	}

	output := ""

	currentNode := rune(input[0])
	count := 1

	for i, r := range input {
		if i == 0 {
			continue
		}
		if r != currentNode {
			if count > 1 {
				output += strconv.Itoa(count)
			}
			output += string(currentNode)
			currentNode = r
			count = 1
		} else {
			count++
		}
	}

	if count > 1 {
		output += strconv.Itoa(count)
	}
	output += string(currentNode)

	return output
}

func RunLengthDecode(input string) string {
	output := ""
	
	digits := []rune{}

	for _, r := range input {
		if unicode.IsDigit(r) {
			digits = append(digits, r)
		} else {
			if len(digits) > 0 {
				n, _ := strconv.Atoi(string(digits))
				output += strings.Repeat(string(r), n)
				digits = digits[:0]
			} else {
				output += string(r)
			}
		}
	}
	
	return output
}
