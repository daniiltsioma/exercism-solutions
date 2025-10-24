package luhn

import (
	"strings"
	"strconv"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return false
	}

	sl := strings.Split(id, "")
	for i := len(sl) - 2; i >= 0; i -= 2 {
		v, err := strconv.Atoi(sl[i])
		if err != nil {
			return false
		}
		v *= 2
		if v > 9 {
			v -= 9
		}
		sl[i] = strconv.Itoa(v)
	}

	sum := 0
	for _, ch := range sl {
		v, err := strconv.Atoi(ch)
		if err != nil {
			return false
		}
		sum += v
	}

	return sum % 10 == 0
}
