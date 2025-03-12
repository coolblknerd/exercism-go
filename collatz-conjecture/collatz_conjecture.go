package collatzconjecture

import "errors"

func CollatzConjecture(n int) (steps int, err error) {
	for steps = 0; n != 1; steps++ {
		if n <= 0 {
			return -1, errors.New("zero or negative value is an error")
		}

		if n%2 == 0 {
			n = n / 2
		} else {
			n = 3*n + 1
		}
	}

	return steps, nil
}
