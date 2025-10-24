package matrix

// import "fmt"

type Pair [2]int

// findLargest finds indexes of largest items.
func findLargest(nums []int) []int {
	indexes := []int{0}
	val := nums[0]

	// don't forget to offset i by 1!
	for i, v := range nums[1:] {
		if v > val {
			indexes = []int{i+1}
			val = v
		} else if v == val {
			indexes = append(indexes, i+1)
		}
	}

	return indexes
}

// findSmallest finds indexes of smallest items.
func findSmallest(nums []int) []int {
	indexes := []int{0}
	val := nums[0]

	// don't forget to offset i by 1!
	for i, v := range nums[1:] {
		if v < val {
			indexes = []int{i+1}
			val = v
		} else if v == val {
			indexes = append(indexes, i+1)
		}
	}

	return indexes
}

func (m *Matrix) Saddle() []Pair {
	pairs := []Pair{}

	if len(*m) == 0 {
		return pairs
	}

	// one-dimensional flat map of trees.
	fm := map[int]bool{}

	// for every tree that is largest in its row,
	// add a map entry.
	for i, row := range m.Rows() {
		for _, r := range findLargest(row) {
			treeIndex := i * len(row) + r
			fm[treeIndex] = true
		}
	}

	// same for every tree that is smallest in its column.
	// return all matches.
	for i, col := range m.Cols() {
		for _, c := range findSmallest(col) {
			treeIndex := c * len((*m)[0]) + i	
			if _, ok := fm[treeIndex]; ok {
				pairs = append(pairs, Pair{c+1, i+1})
			}
		}
	}

	return pairs
}