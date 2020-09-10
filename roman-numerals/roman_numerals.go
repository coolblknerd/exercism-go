package romannumerals

type RomanNumeral map[int]string

var rn = RomanNumeral{
	1: "I",
	5: "V",
	10: "X",
	50: "L",
	100: "C",
	500: "D",
	1000: "M",
}

func ToRomanNumeral(i int) string {
	var numeral []string

	if i < 5 {
		
	} else if i >= 5 && i < 10 {

	} else if i >= 10 && i < 50 {

	} else if i >= 50 && i < 100 {

	} else if i >= 100 && i < 500 {

	} else if i >= 500 && i <= 1000 {

	}
}