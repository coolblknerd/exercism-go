package isogram

import "unicode"

// IsIsogram is a function that determines if word is a isogram
func IsIsogram(word string) bool {
	seen := make(map[rune]bool)
	for _, r := range word {
		if !unicode.IsLetter(r) {
			continue
		}
		lc := unicode.ToLower(r)
		if seen[lc] {
			return false
		}
		seen[lc] = true
	}
	return true
}
