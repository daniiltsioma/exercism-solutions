package minesweeper

// Annotate returns an annotated board
func Annotate(board []string) []string {
	neighbors := [][2]int{
		{-1, -1}, {-1, 0}, {-1,1}, {0, 1},
		{1, 1}, {1, 0}, {1, -1}, {0, -1},
	}

	// annotated board
	ab := []string{}
	for r, row := range board {
		ar := []rune{}
		for c, char := range row {
			if char == '*' {
				ar = append(ar, char)
			} else {
				count := 0
				for _, n := range neighbors {
					y, x := r + n[0], c + n[1]
					if y < 0 || y >= len(board) || x < 0 || x >= len(row) {
						continue
					}
					if board[y][x] == byte('*') {
						count++
					}
				}
				if count > 0 {
					ar = append(ar, rune(48 + count))
				} else {
					ar = append(ar, ' ')
				}
			}
		}
		ab = append(ab, string(ar))
	}

	return ab
}
