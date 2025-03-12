package luhn

import (
	"regexp"
	"strconv"
	"strings"
)

func Valid(id string) (valid bool) {
	sum := 0

	id = strings.ReplaceAll(id, " ", "")

	if len(id) <= 1 {
		return valid
	}

	re := regexp.MustCompile(`\D+`)
	if re.MatchString(id) {
		return valid
	}

	secondDigit := len(id)%2 == 0

	for _, r := range id {
		v, err := strconv.Atoi(string(r))
		if err != nil {
			return valid
		}

		if secondDigit {
			v *= 2
			if v > 9 {
				v -= 9
			}
		}

		sum += v
		secondDigit = !secondDigit
	}

	return sum%10 == 0
}
