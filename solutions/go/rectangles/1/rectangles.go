package rectangles

func Count(diagram []string) int {
	count := 0

	for r, row := range diagram {
		for c, tl := range row {
			if tl == '+' { // found top left plus
				for i:=r+1; i < len(diagram); i++ {
					if diagram[i][c] == '+' { // found bottom left
						for j:=c+1; j < len(row); j++ {
							// top right, bottom right
							tr, br := diagram[r][j], diagram[i][j]

							if tr == '+' && br == '+' {
								valid := true
								for k:=r+1; k<i; k++ {
									if diagram[k][j] != '|' && diagram[k][j] != '+' {
										valid = false
										break
									}
								}
								if valid {
									count++
								}
								continue
							} else if tr == ' ' || br == ' ' {
								break
							}
						}
						continue
					}
					if diagram[i][c] != '|' {
						break
					}
				}
			}
		}
	}

	return count
}
