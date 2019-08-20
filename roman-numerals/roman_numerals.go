package romannumerals

import (
	"errors"
	"strings"
)

const (
	I = "I"
	V = "V"
	X = "X"
	L = "L"
	C = "C"
	D = "D"
	M = "M"
)

// ToRomanNumeral takes a number and converts it to a roman numeral
func ToRomanNumeral(number int) (ans string, err error) {
	// This initiates an empty slice
	numerals := []string{}

	for i := 1; i <= number; i++ {
		if number < 5 {
			if number == 4 {
				numerals = append(numerals, "IV")
				break
			} else {
				numerals = append(numerals, I)
			}
		}

		if number >= 5 && number < 10 {
			if number == 5 {
				numerals = append(numerals, V)
				break
			} else {

			}
		}
	}

	ans = strings.Join(numerals, "")

	if ans == "" {
		err = errors.New("Something fucked up")
	}

	return ans, err
}
