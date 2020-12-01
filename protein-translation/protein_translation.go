package protein

import (
	"errors"
)

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("Invalid Codon")

var aminos = map[string]string{
	"AUG": "Methionine",
	"UUU": "Phenylalanine",
	"UUC": "Phenylalanine",
	"UUA": "Leucine",
	"UUG": "Leucine",
	"UCU": "Serine",
	"UCC": "Serine",
	"UCA": "Serine",
	"UCG": "Serine",
	"UAU": "Tyrosine",
	"UAC": "Tyrosine",
	"UGU": "Cysteine",
	"UGC": "Cysteine",
	"UGG": "Tryptophan",
	"UAA": "STOP",
	"UAG": "STOP",
	"UGA": "STOP",
}

func FromCodon(input string) (string, error) {
	value, ok := aminos[input]

	if !ok {
		return "", ErrInvalidBase
	} else if value == "STOP" {
		return "", ErrStop
	}

	return value, nil
}

func FromRNA(input string) ([]string, error) {
	pro := []string{}

	for i := 0; i < len(input); i += 3 {
		c := input[i : i+3]

		v, ok := aminos[c]
		if !ok {
			return pro, ErrInvalidBase
		} else if v == "STOP" {
			break
		}

		for codon, protein := range aminos {
			if c == codon {
				pro = append(pro, protein)
			}
		}
	}
	return pro, nil
}
