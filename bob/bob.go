package bob

import (
	"strings"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	// str := strings.Split(remark, " ")
	// firstWord := str[0]
	// firstElement := []rune(firstWord)[0]
	// lastElement := string(remark[len(remark)-1])

	str := strings.TrimSpace(remark)
	// fmt.Println("-------")
	// fmt.Println("Slice of Words: ", str)
	// fmt.Println("First Word: ", firstWord)
	// fmt.Println("First Element", firstElement)
	// fmt.Println("Last Element: ", lastElement)
	// fmt.Println("Length of Sentence: ", len(str))

	switch {
	default:
		return "Whatever."
	case str == "":
		return "Fine. Be That Way"
	case string(str[len(str)-1]) == "?":
		if str == strings.ToUpper(str) {
			return "Calm down, I know what I'm doing!"
		}
		return "Sure."
	case str == strings.ToUpper(str):
		return "Whoa, chill out!"
	}
}
