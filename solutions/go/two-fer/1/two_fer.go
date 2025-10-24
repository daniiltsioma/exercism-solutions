// Package twofer contains tools for
// addressing the receiver of a complimentary cookie.
package twofer

import "fmt"

// Sharewith returns an appropriate message for the receiver of the cookie.
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	}

	return fmt.Sprintf("One for %s, one for me.", name)
}
