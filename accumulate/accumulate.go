package accumulate

type convert func(string) string

func Accumulate(c []string, fn convert) (collection []string) {
	for _, v := range c {
		collection = append(collection, fn(v))
	}
	return
}
