package romannumerals

import "errors"

func ToRomanNumeral(input int) (result string, err error) {
	romanMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}

	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	if input <= 0 || input >= 4000 {
		return "", errors.New("input is out of range")
	}

	for _, value := range values {
		for input >= value {
			result += romanMap[value]
			input -= value
		}
	}

	return result, nil
}
