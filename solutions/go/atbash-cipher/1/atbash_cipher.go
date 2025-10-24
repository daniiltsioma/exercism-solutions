package atbash

import (
	"unicode"
	"strings"
)

func Atbash(s string) string {
	s = strings.ToLower(s)	

	res := []rune{}
	groupLen := 0

	for _, r := range s {
		if !(unicode.IsLetter(r) || unicode.IsDigit(r)) {
			continue
		}

		if groupLen == 5 {
			res = append(res, ' ')	
			groupLen = 0
		}

		if unicode.IsDigit(r) {
			res = append(res, r)
		} else {
			res = append(res, rune(122 - (int(r) - 97)))
		}

		groupLen++
	}

	return string(res)
}
