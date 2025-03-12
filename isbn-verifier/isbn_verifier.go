package isbn

import "strings"

func IsValidISBN(isbn string) (valid bool) {
	isbn = strings.ReplaceAll(isbn, "-", "")
	if len(isbn) != 10 {
		return
	}

	var sum int
	for i, c := range isbn {
		if i == 9 && c == 'X' {
			sum += 10
			break
		}

		if c < '0' || c > '9' {
			return
		}

		sum += int(c-'0') * (10 - i)
	}

	return sum%11 == 0
}
