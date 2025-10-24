package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("Cannot compare strands of different length.")
	}

	dist := 0

	for i := range len(a) {
		if a[i] != b[i] {
			dist++
		}
	}

	return dist, nil
}
