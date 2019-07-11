package hamming

import (
	"errors"
)

// Distance calculates the Hamming distance between two stands of DNA. (Always end comments with periods)
func Distance(a, b string) (int, error) {
	var dist int

	if len(a) != len(b) {
		return dist, errors.New("lengths are not the same") // It's suggested to not use caps or punctuations in errors to leave breadcrumb effect
	}

	ra, rb := []rune(a), []rune(b)

	for i := range ra {
		if rb[i] != ra[i] {
			dist++
		}
	}

	return dist, nil
}
