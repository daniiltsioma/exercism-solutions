package reverse

func Reverse(input string) string {
	runes := []rune(input)
	runesRev := []rune{}

	for i := len(runes) - 1; i >= 0; i-- {
		runesRev = append(runesRev, runes[i])
	}
	
	return string(runesRev)
}
