package reverse

func Reverse(s string) (r string) {
	for _, l := range s {
		r = string(l) + r
	}
	return
}
