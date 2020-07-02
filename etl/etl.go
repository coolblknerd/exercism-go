package etl

import (
	"strings"
)

// Transform converts a type
func Transform(m map[int][]string) map[string]int {
	output := make(map[string]int)
	for k, v := range m {
		for _, value := range v {
			output[strings.ToLower(value)] = k
		}
	}
	return output
}
