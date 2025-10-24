package isogram

import "strings"

func IsIsogram(word string) bool {
	word = strings.ToLower(word)
	word = strings.ReplaceAll(word, " ", "")
	word = strings.ReplaceAll(word, "-", "")

	chars := map[rune]int{}

	for _, ch := range word {
		_, exists := chars[ch]; if exists {
			return false
		}
		chars[ch] = 1
	}

	return true
}
