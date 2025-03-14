package resistorcolorduo

import (
	"strconv"
)

// Colors returns the list of all colors.
func Colors() []string {
	return []string{
		"black", "brown", "red", "orange", "yellow",
		"green", "blue", "violet", "grey", "white",
	}
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	for i, c := range Colors() {
		if c == color {
			return i
		}
	}

	return -1
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) (resistance int) {
	value := ""

	for i := 0; i < 2; i++ {
		if ColorCode(colors[i]) == -1 {
			return -1
		}

		value += strconv.Itoa(ColorCode(colors[i]))
	}

	resistance, _ = strconv.Atoi(value)

	return resistance
}
