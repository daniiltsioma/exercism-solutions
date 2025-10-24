package brackets

func Bracket(input string) bool {
	stack := []rune{}

	for _, r := range input {
		switch r {
		case '}':
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ']':
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		case ')':
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}
			stack = stack[:len(stack)-1]
		case '{', '[', '(':
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}
