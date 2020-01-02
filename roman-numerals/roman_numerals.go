package romannumerals

import (
	"errors"
	"strings"
)

type romanNumeral struct {
	value  int
	symbol string
}

// ToRomanNumeral takes a number and converts it to a roman numeral
func ToRomanNumeral(number int) (ans string, err error) {
	if number <= 0 || number > 3000 {
		return "", errors.New("Number is too large")
	}
	// This initiates an empty slice
	numerals := []string{}

	romanNumerals := []romanNumeral{
		{1, "I"},
		{4, "IV"},
		{5, "V"},
		{9, "IX"},
		{10, "X"},
		{40, "XL"},
		{50, "L"},
		{90, "XC"},
		{100, "C"},
		{400, "CD"},
		{500, "D"},
		{900, "CM"},
		{1000, "M"},
	}

	for _, numeral := range romanNumerals {
		if number == numeral.value {
			numerals = append(numerals, numeral.symbol)
		}

		if number < numeral.value {
			numerals = append(numerals)
		}
	}

	return strings.Join(numerals, ""), nil
}
