package wordcount

import (
	"regexp"
	"strings"
)

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	freq := Frequency{}
	re := regexp.MustCompile(`[:!?,;:\s\.]`)
	words := re.Split(phrase, -1)

	for _, w := range words {
		re = regexp.MustCompile(`[&@$%^]`)
		w = string(re.ReplaceAll([]byte(w), []byte{}))

		w = strings.Trim(w, "' ")
		if len(w) == 0 {
			continue
		}

		w = strings.ToLower(w)

		_, ok := freq[w]; if !ok {
			freq[w] = 1
		} else {
			freq[w]++
		}
	}

	return freq
}
