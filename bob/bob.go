package bob

import "strings"

// Hey should have a comment documenting it.
func Hey(remark string) string {
	lastElement := string(remark[len(remark)-1])
	if remark == strings.ToUpper(remark) && lastElement != "?" {
		return "Whoa, chill out!"
	} else if lastElement == "?" && remark != strings.ToUpper(remark) {
		return "Sure."
	} else if lastElement == "?" && remark == strings.ToUpper(remark) {
		return "Calm down, I know what I'm doing!"
	}
	return "Whatever."
}
