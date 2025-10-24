// Package bob helps predict Bob's next phrase.
package bob

import (
	"unicode"
	"strings"
)

// Hey determines how Bob is going to respond to the remark.
func Hey(remark string) string {
	if len(strings.Trim(remark, " \n\t\r")) == 0 {
		return "Fine. Be that way!"
	}

	remark = strings.Trim(remark, " ")

	isQuestion := remark[len(remark) - 1] == '?'

	hasLetters := false
	isUpper := true

	for _, ch := range remark {
		if unicode.IsLetter(ch) {
			hasLetters = true
		}
		if unicode.IsLetter(ch) && unicode.IsLower(ch) {
			isUpper = false
		}		
	}

	if isQuestion && isUpper && hasLetters {
		return "Calm down, I know what I'm doing!"
	} else if isQuestion {
		return "Sure."
	} else if isUpper && hasLetters {
		return "Whoa, chill out!"
	}

	return "Whatever."
}
