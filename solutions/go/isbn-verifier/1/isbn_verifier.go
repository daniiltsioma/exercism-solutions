package isbn

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.Replace(isbn, "-", "", -1)
	if len(isbn) != 10 {
		return false
	}

	sum := 0
	for i, d := range isbn {
		var v int
		if d == 'X' {
			if i == 9 {
				v = 10
			} else {
				return false
			}
		} else {
			v = int(d - '0')
			if v < 0 || v > 9 {
				return false
			}
		}
		sum += v * (10 - i)
	}

	return sum % 11 == 0
}
