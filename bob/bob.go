package bob

import (
	"regexp"
	"unicode"
)

// Hey is a function to test if a remark is as expected
func Hey(remark string) string {
	last := remark[len(remark)-1:]
	match, _ := regexp.MatchString("[.!?0-9]", last)

	if match == true {
		switch last {
		case "!":
			if !isUppercase(remark) {
				return "Whatever."
			}
			return "Whoa, chill out!"
		case "?":
			if isUppercase(remark) {
				return "Calm down, I know what I'm doing!"
			}
			return "Sure."
		default:
			return "Whatever."
		}
	}

	if isUppercase(remark) {
		return "Whoa, chill out!"
	}

	return "Nothing found yet."
}

func isUppercase(word string) bool {
	count := 0
	for _, r := range word {
		if unicode.IsNumber(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) || unicode.IsSpace(r) {
			count++
			continue
		}
		match, _ := regexp.MatchString("[A-Z]", string(r))
		if !match {
			return false
		}
	}

	if count == len(word) {
		return false
	}

	return true
}
