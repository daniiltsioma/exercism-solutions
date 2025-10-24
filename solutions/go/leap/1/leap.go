// Package leap helps determine if a given year
// is a leap year or not.
package leap

// IsLeapYear returns a boolean telling whether
// the year is a leap year or not.
func IsLeapYear(year int) bool {
	if year % 100 == 0 {
		return year % 400 == 0
	}

	return year % 4 == 0
}
