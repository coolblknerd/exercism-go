package wordcount

import (
	"regexp"
	"strings"
)

// Frequency is the amount a word shows up in a sentence
type Frequency map[string]int

// Returns a Frequency of words in a sentence
func WordCount(input string) Frequency {
	reg, _ := regexp.Compile("[^a-zA-Z0-9']+")

	words := reg.Split(input, -1)
	f := Frequency{}
	for _, v := range words {
		lc := strings.ToLower(v)
		f[strings.Trim(lc, "'")] += 1
	}

	if _, ok := f[""]; ok {
		delete(f, "")
	}

	return f
}
