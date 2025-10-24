package prime

func Factors(n int64) []int64 {
	divisors := []int64{}
	
	div := int64(2)

	for n > 1 {
		if n % div == 0 {
			divisors = append(divisors, div)
			n /= div
		} else {
			div++
		}
	}

	return divisors
}
