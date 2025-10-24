package pangram

import "strings"

func IsPangram(input string) bool {
	input = strings.ToLower(input)
	apps := map[rune]int{}

	for _, ch := range input {
		apps[ch] = 1
	}

	// go through lowercase ASCII
	for i := 97; i < 123; i++ {
		_, ok := apps[rune(i)]; if !ok {
			return false
		}
	}

	return true
}
