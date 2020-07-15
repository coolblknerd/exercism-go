package isogram

import "unicode"

// IsIsogram is a function that determines if word is a isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, r := range word {
		lc := unicode.ToLower(r)
		if seen[lc] {
			return false
		} else if !unicode.IsLetter(r) {
			continue
		}
		seen[lc] = true
	}
	return true
}
