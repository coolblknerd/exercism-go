package isogram

import (
	"strings"
	"unicode"
)

// IsIsogram is a function that determines if word is a isogram
func IsIsogram(word string) bool {
	var l []string
	for _, rune := range word {
		if strings.ContainsRune(strings.Join(l, ""), unicode.ToLower(rune)) {
			return false
		} else if string(rune) == "-" || string(rune) == " " {
			continue
		} else {
			lc := unicode.ToLower(rune)
			l = append(l, string(lc))
		}
	}
	return true
}
