package railfence

func Encode(message string, rails int) string {
	lines := make([][]rune, rails)

	row := 0
	modifier := 1

	for _, ch := range message {
		lines[row] = append(lines[row], ch)
		if row == 0 {
			modifier = 1
		} else if row == rails - 1 {
			modifier = -1
		}
		row += modifier
	}

	output := ""

	for _, l := range lines {
		output += string(l)
	}

	return output
}

func Decode(message string, rails int) string {
	row := 0
	modifier := 1
	
	// calculate row lengths 
	rowLengths := make([]int, rails)

	for _ = range message {
		rowLengths[row]++
		if row == 0 {
			modifier = 1
		} else if row == rails - 1 {
			modifier = -1
		}
		row += modifier
	}

	// split into lines
	lines := []string{}
	offset := 0
	for	_, l := range rowLengths {
		lines = append(lines, message[offset:offset+l])
		offset += l
	} 

	decoded := []byte{}

	// decode
	row = 0
	modifier = 1
	cols := map[int]int{} // keep track of current columns

	for _ = range message {
		c := cols[row]
		decoded = append(decoded, lines[row][c])
		if row == 0 {
			modifier = 1
		} else if row == rails - 1 {
			modifier = -1
		}
		cols[row]++
		row += modifier
	}

	return string(decoded)
}
