package anagram

import "strings"

func buildMap(word string) map[rune]int {
	word = strings.ToLower(word)
	m := map[rune]int{}
	for _, ch := range word {
		_, ok := m[ch]; if !ok {
			m[ch] = 1
		} else {
			m[ch]++
		}
	}

	return m
}

func compareMaps(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		v, ok := m2[k]; if !ok {
			return false
		} else if m1[k] != v {
			return false
		}
	}
	return true
}

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	
	subjectMap := buildMap(subject)

	for _, cand := range candidates {
		candMap := buildMap(cand)
		if (strings.ToLower(subject) != strings.ToLower(cand)) && 
			compareMaps(subjectMap, candMap) {
			anagrams = append(anagrams, cand)
		}
	}

	return anagrams
}
