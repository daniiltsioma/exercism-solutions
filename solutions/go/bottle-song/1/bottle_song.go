package bottlesong

import (
	"unicode"
	"fmt"
)


func btl(b int) string {
	if b == 1 {
		return "bottle"
	}
	return "bottles"
}

func cap(s string) string {
	runes := []rune(s)
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

func Recite(startBottles, takeDown int) []string {
	nums := []string{"no", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

	res := []string{}

	for i := 0; i < takeDown; i++ {
		if i > 0 {
			res = append(res, "")
		}
		n := nums[startBottles]
		b := btl(startBottles)
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", cap(n), b))
		res = append(res, fmt.Sprintf("%s green %s hanging on the wall,", cap(n), b))
		res = append(res, fmt.Sprintf("And if one green bottle should accidentally fall,"))

		startBottles--
		n = nums[startBottles]
		b = btl(startBottles)
		res = append(res, fmt.Sprintf("There'll be %s green %s hanging on the wall.", n, b))
	}

	return res
}
