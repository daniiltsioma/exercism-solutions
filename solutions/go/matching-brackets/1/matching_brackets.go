package brackets

func Bracket(input string) bool {
	brackets := []rune{}

	for _, r := range input {
		switch r {
		case '}': 
			if len(brackets) == 0 || brackets[len(brackets)-1] != '{' {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		case ')': 
			if len(brackets) == 0 || brackets[len(brackets)-1] != '(' {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		case ']': 
			if len(brackets) == 0 || brackets[len(brackets)-1] != '[' {
				return false
			}
			brackets = brackets[:len(brackets)-1]
		case '(', '{', '[':
			brackets = append(brackets, r)
		}
	}

	return len(brackets) == 0
}
