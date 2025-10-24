// Package proverb generates proverbs from user input.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import "fmt"

// Proverb takes a slice of rhymes and returns a slice of proverbs
func Proverb(rhyme []string) []string {
	proverbs := []string{}

	for i, r := range rhyme {
		if i < len(rhyme) - 1 {
			proverbs = append(proverbs, fmt.Sprintf("For want of a %s the %s was lost.", r, rhyme[i+1]))
		}
		if i == len(rhyme) - 1 {
			proverbs = append(proverbs, fmt.Sprintf("And all for the want of a %s.", rhyme[0]))
		}
	}

	return proverbs
}
