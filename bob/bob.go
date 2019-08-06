package bob

import (
	"fmt"
	"strconv"
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	lastElement := string(remark[len(remark)-1])
	ans := "Whatever."
	fmt.Println(lastElement)

	switch lastElement {
	default:
		if _, err := strconv.Atoi(lastElement); err == nil {
			ans = "Whatever."
		} else if strings.TrimSpace(remark) == "" {
			ans = "Fine. Be that way!"
		} else if remark == strings.ToUpper(remark) {
			ans = "Whoa, chill out!"
		}
	case "?":
		if remark == strings.ToUpper(remark) && {
			ans = "Calm down, I know what I'm doing!"
		} else {
			ans = "Sure."
		}
	}

	return ans
}
