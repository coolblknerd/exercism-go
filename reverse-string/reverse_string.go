package reverse

func Reverse(s string) (r string) {
	n := len(s)
	rev := make([]rune, n)
	for _, v := range s {
		n--
		rev[n] = v
	}
	return string(rev[n:])
}
