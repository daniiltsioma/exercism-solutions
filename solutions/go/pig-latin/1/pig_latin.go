package piglatin

import (
	"strings"
)

func Sentence(sentence string) string {
	words := strings.Split(sentence, " ")
	
	wordsEncrypted := []string{}

	vowels := map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	rotated := false

	for _, word := range words {
		_, wordStartsWithVowel := vowels[rune(word[0])]
		// rule 1
		if wordStartsWithVowel || word[:2] == "xr" || word[:2] == "yt" {
			word += "ay"
			wordsEncrypted = append(wordsEncrypted, word)
			continue
		}

		// move consonants to the end
		for {
			if word[:2] == "qu" {
				word = word[2:] + "qu"
				rotated = true
			} else if _, ok := vowels[rune(word[0])]; !ok {
				if word[0] == 'y' && rotated {
					break
				}
				word = word[1:] + string(word[0])
				rotated = true
			} else {
				break
			}
		}
		word += "ay"
		wordsEncrypted = append(wordsEncrypted, word)
		continue
	}

	res := ""
	for i, w := range wordsEncrypted {
		res += w
		if i < len(words) - 1 {
			res += " "
		}
	}

	return res
}
