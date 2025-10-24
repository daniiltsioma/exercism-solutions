package series

func All(n int, s string) []string {
	substrings := []string{}

	for i := 0; i < len(s) - n + 1; i++ {
		substrings = append(substrings, s[i:i+n])
	}

	return substrings
}

func UnsafeFirst(n int, s string) string {
	return s[:n]
}

func First(n int, s string) (first string, ok bool) {
	if len(s) < n {
		return "", false
	}
	return s[:n], true
}