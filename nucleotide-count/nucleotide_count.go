package dna

import (
	"errors"
	"strings"
)

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram struct {
	A int
	G int
	C int
	T int
}

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

// Counts generates a histogram of valid nucleotides in the given DNA.
// Returns an error if d contains an invalid nucleotide.
///
// Counts is a method on the DNA type. A method is a function with a special receiver argument.
// The receiver appears in its own argument list between the func keyword and the method name.
// Here, the Counts method has a receiver of type DNA named d.
func (d DNA) Counts() (Histogram, error) {
	var h Histogram
	var invalid error
	dna := string(d)
	nucleotides := strings.Split(dna, "")

	for _, n := range nucleotides {
		if n == "A" {
			h.A++
		} else if n == "G" {
			h.G++
		} else if n == "C" {
			h.C++
		} else if n == "T" {
			h.T++
		} else {
			invalid = errors.New("Strand has invalid nucleotides")
		}
	}

	return h, invalid

}
