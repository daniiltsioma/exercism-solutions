package summultiples

func SumMultiples(limit int, divisors ...int) int {
	mmap := map[int]bool{}
	msl := []int{}

	for _, d := range divisors {
		if d == 0 {
			continue
		}
		for i := d; i < limit; i = i + d {
			_, ok := mmap[i]; if !ok {
				msl = append(msl, i)
			}
			mmap[i] = true
		}
	}

	sum := 0
	for _, v := range msl {
		sum += v
	}

	return sum
}
