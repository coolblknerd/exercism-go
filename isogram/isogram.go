package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram is a function that determines if word is a isogram
func IsIsogram(word string) bool {
	var l []string
	for _, r := range word {
		if strings.ContainsRune(strings.Join(l, ""), unicode.ToLower(r)) {
			return false
		} else if string(r) == "-" || string(r) == " " {
			continue
		} else {
			lc := unicode.ToLower(r)
			l = append(l, string(lc))
		}
	}
	return true
}
