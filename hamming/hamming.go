package hamming

import (
	"errors"
	"strings"
)

// Distance calculates the Hamming distance between two stands of DNA
func Distance(a, b string) (int, error) {
	var err error
	var dist int

	if len(a) != len(b) {
		err = errors.New("Lengths are not the same")
	} else if a != b {
		arrA := strings.Split(a, "")
		arrB := strings.Split(b, "")

		for i, value := range arrA {
			if arrB[i] != value {
				dist++
			}
		}
	}

	return dist, err
}
