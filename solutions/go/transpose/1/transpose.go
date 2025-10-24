package transpose

func Transpose(input []string) []string {
	if len(input) == 0 {
		return []string{}
	}

	numRows := 0
	for _, row := range input {
		if len(row) > numRows {
			numRows = len(row)
		}
	}	

	lines := make([][]byte, numRows)

	trueLengths := map[int]int{}

	for c := range numRows {
		for _, row := range input {
			if c < len(row) {
				lines[c] = append(lines[c], row[c])
				trueLengths[c] = len(lines[c])
			} else {
				lines[c] = append(lines[c], ' ')
			}
		}
	}

	res := []string{}

	for i, line := range lines {
		res = append(res, string(line[:trueLengths[i]]))
	}

	return res
}
