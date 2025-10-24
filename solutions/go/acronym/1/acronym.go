// Package acronym contains tools to generate acronyms.
package acronym

import (
	"unicode"
	"regexp"
	"strings"
)

// Abbreviate returns an acronym for a given string.
func Abbreviate(s string) string {
	abbr := []rune{}

	// replace all unnecessary characters.
	re := regexp.MustCompile(`[,_']`)
	s = string(re.ReplaceAll([]byte(s), []byte("")))

	// trim spaces around hyphens.
	s = strings.ReplaceAll(s, " - ", "-")

	// split by whitespaces and hyphens.
	re = regexp.MustCompile(`[-\s]`)
	parts := re.Split(s, -1)


	for _, p := range parts {
		abbr = append(abbr, unicode.ToUpper(rune(p[0])))
	}

	return string(abbr)
}
