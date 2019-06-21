package hamming

import (
	"errors"
)

// Distance calculates the Hamming distance between two stands of DNA. (Always end comments with periods)
func Distance(a, b string) (int, error) {
	var err error
	var dist int

	if len(a) != len(b) {
		return dist, errors.New("lengths are not the same") // It's suggested to not use caps or punctuations in errors to leave breadcrumb effect
	} else if a != b {
		for i, value := range a {
			if rune(b[i]) != value {
				dist++
			}
		}
	}

	return dist, err
}
