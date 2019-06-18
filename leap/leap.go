// Package leap should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package leap

// IsLeapYear should have a comment documenting it.
func IsLeapYear(year int) bool {
	everyYear := year % 4
	everyCentury := year % 100
	everyFourHundredYears := year % 400

	if everyYear != 0 || everyCentury == 0 && everyFourHundredYears != 0 {
		return false
	}

	return true
}
