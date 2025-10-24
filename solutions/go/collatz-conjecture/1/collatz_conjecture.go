package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("Unable to prove the conjecture for a non-positive number.")
	}

	if n == 1 {
		return 0, nil
	}

	// odd
	if n % 2 == 1 {
		res, _ := CollatzConjecture(3 * n + 1)
		return 1 + res, nil
	}

	// even
	res, _ := CollatzConjecture(n / 2)
	return 1 + res, nil
}
