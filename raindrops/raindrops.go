package raindrops

import (
	"strconv"
	"strings"
)

// Convert should take a number and turn it to a string depending on the factorial of the given parameter
func Convert(value int) string {
	var drops []string

	if value%3 == 0 {
		drops = append(drops, "Pling")
	}

	if value%5 == 0 {
		drops = append(drops, "Plang")
	}

	if value%7 == 0 {
		drops = append(drops, "Plong")
	}

	if len(drops) >= 1 {
		return strings.Join(drops, "")
	}

	return strconv.Itoa(value)
}
