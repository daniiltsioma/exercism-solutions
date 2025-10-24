package prime

import "errors"

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("can't return the nth prime for n < 1")
	}
	primes := []int{}

	i := 2
	
	for len(primes) < n {
		divisible := false
		for _, p := range primes {
			if i % p == 0 {
				divisible = true
				break
			}
		}
		if !divisible {
			primes = append(primes, i)
		}
		i++
	}

	return primes[n-1], nil
}
