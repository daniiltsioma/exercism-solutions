package sieve

func Sieve(limit int) []int {
	m := map[int]bool{}
	for i := 2; i <= limit; i++ {
		m[i] = false
	}

	for i := 2; i <= limit; i++ {
		if m[i] {
			continue
		}
		for j := i * 2; j <= limit; j += i {
			m[j] = true
		}
	}

	res := []int{}

	for i := 2; i <= limit; i++ {
		if !m[i] {
			res = append(res, i)
		}
	}

	return res
}
