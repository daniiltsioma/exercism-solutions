package wordy

import (
	"strings"
	"strconv"
)

func Answer(question string) (int, bool) {
	// replace math operations for simplicity.
	question = strings.ReplaceAll(question, "plus", "+")
	question = strings.ReplaceAll(question, "minus", "-")
	question = strings.ReplaceAll(question, "multiplied by", "*")
	question = strings.ReplaceAll(question, "divided by", "/")
	
	// replace the question mark at the end.
	question = strings.Replace(question, "?", "", 1)

	// split into parts.
	parts := strings.Split(question, " ")

	// non-math questions are not allowed.
	if parts[0] != "What" || parts[1] != "is" {
		return 0, false
	}

	// drop the first two words.
	parts = parts[2:]

	// reject if there are no operands.
	if len(parts) == 0 {
		return 0, false
	}

	// reduce until we have the last item left.
	for len(parts) > 1 {
		// can't have even number of parts.
		if len(parts) % 2 == 0 {
			return 0, false
		}
		// extract left operand, operation, and right operand as strings.
		ls, ops, rs := parts[0], parts[1], parts[2]	
		// operands must be digits.
		l, err := strconv.Atoi(ls)
		if err != nil {
			return 0, false
		}
		r, err := strconv.Atoi(rs)
		if err != nil {
			return 0, false
		}
		// op must be a valid operation.
		if len(ops) > 1 || !strings.ContainsRune("+-*/", rune(ops[0])) {
			return 0, false
		}
		op := ops[0]
		// calculate the result.
		var res int
		switch op {
		case '+':
			res = l + r
		case '-':
			res = l - r
		case '*':
			res = l * r
		case '/':
			res = l / r
		default:
			return 0, false
		}
		resStr := strconv.Itoa(res)
		// reduce the parts slice.
		if len(parts) > 3 {
			parts = append([]string{resStr}, parts[3:]...)
		} else {
			parts = []string{resStr}
		}

	}
	
	res, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, false
	}

	return res, true
}