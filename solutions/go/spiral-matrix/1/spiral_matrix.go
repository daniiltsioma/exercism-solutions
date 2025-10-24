package spiralmatrix

type Direction int

const (
	East Direction = iota
	South
	West
	North
)

func SpiralMatrix(size int) [][]int {
	matrix := make([][]int, size)
	for i := 0; i < size; i++ {
		matrix[i] = make([]int, size)
	}

	var direction Direction = 0

	top := 0
	left := 0

	right := size - 1
	bottom := size - 1

	r, c := 0, 0

	for v := 1; v <= size*size; v++ {
		matrix[r][c] = v

		switch direction {
		case East:
			if c < right { // keep going
				c++
			} else { // turn and shrink
				direction = South
				r++
				top++
			}
		case South:
			if r < bottom { // keep going
				r++
			} else { // turn and shrink
				direction = West
				c--
				right--
			}
		case West:
			if c > left { // keep going
				c--
			} else { // turn and shrink
				direction = North
				r--
				bottom--
			}
		case North:
			if r > top { // keep going
				r--
			} else { // turn and shrink
				direction = East
				c++
				left++
			}
		}
	}

	return matrix
}
