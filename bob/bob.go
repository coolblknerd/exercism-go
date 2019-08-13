package bob

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	str := strings.Split(remark, " ")
	firstWord := str[0]
	firstElement := []rune(firstWord)[0]
	lastElement := string(remark[len(remark)-1])

	ans := "Whatever."
	fmt.Println("-------")
	fmt.Println("Slice of Words: ", str)
	fmt.Println("First Word: ", firstWord)
	fmt.Println("First Element", firstElement)
	fmt.Println("Last Element: ", lastElement)
	fmt.Println("Length of Sentence: ", len(str))

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
		if firstWord == strings.ToUpper(firstWord) && unicode.IsLetter(firstElement) {
			ans = "Calm down, I know what I'm doing!"
		} else {
			ans = "Sure."
		}
	}

	return ans
}
