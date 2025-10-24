package pascal

func Triangle(n int) [][]int {
	if n < 1 {
		return [][]int{}
	}

	rows := make([][]int, n)
	rows[0] = []int{1}

	for i := 1; i < n; i++ {
		row := make([]int, i+1)
		row[0] = 1
		row[i] = 1

		for j := 1; j < i; j++ {
			row[j] = rows[i-1][j-1] + rows[i-1][j]
		}
		rows[i] = row
	}

	return rows
}