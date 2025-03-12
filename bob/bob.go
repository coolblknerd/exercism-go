// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey is a function that returns a string based on the input string
func Hey(remark string) string {
	remark = strings.TrimSpace(remark)

	if len(remark) == 0 {
		return "Fine. Be that way!"
	}

	if strings.HasSuffix(remark, "?") {
		if IsUpperCase(remark) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	}

	if IsUpperCase(remark) {
		return "Whoa, chill out!"
	}

	return "Whatever."
}

func IsUpperCase(s string) bool {
	lettersPresent := false

	for _, r := range s {
		if unicode.IsLetter(r) {
			lettersPresent = true

			if !unicode.IsUpper(r) {
				return false
			}
		}
	}

	return lettersPresent
}
