package binarysearch

func SearchInts(list []int, key int) int {
	l := 0
	r := len(list) - 1

	for l <= r {
		m := (l + r) / 2

		if list[m] < key {
			l = m + 1
		} else if key < list[m] {
			r = m - 1
		} else {
			return m
		}
	}

	return -1
}
