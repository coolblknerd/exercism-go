package bob

import (
	"strings"
)

// Hey is a function to test if a remark is as expected
func Hey(remark string) string {
	str := strings.Split(remark, " ")

	switch str {
	default:
		return "Whatever."

	}
}
