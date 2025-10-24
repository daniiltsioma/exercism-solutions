package logs

import "unicode/utf8"

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, c := range log {
		if c == 0x2757 {
			return "recommendation"
		} else if c == 0x1F50D {
			return "search"
		} else if c == 0x2600 {
			return "weather"
		}
	}
	return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	newLog := []int32{}

	for _, c := range log {
		if c == oldRune {
			newLog = append(newLog, newRune)
		} else {
			newLog = append(newLog, c)
		}
	}

	return string(newLog)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
