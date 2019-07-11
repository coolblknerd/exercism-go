package raindrops

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

// Convert should take a number and turn it to a string depending on the factorial of the given parameter
func Convert(value int) string {

	var drops []string
	var keys []int

	factors := map[int]string{
		3: "Pling",
		5: "Plang",
		7: "Plong",
	}

	// This will sort the keys
	for k := range factors {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	for _, k := range keys {
		if value%k == 0 {
			fmt.Printf("%d just returned 0 for factor of %d\n", value, k)
			drops = append(drops, factors[k])
		}
	}

	if len(drops) >= 1 {
		return strings.Join(drops, "")
	}
	return strconv.Itoa(value)
}
