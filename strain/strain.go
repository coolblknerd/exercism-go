package strain

type Ints []int
type Lists [][]int
type Strings []string

type convert func(int) bool
type stringConvert func(string) bool
type listsConvert func(l []int) bool

func (i Ints) Keep(fn convert) Ints {
	var ans Ints
	for _, value := range i {
		if fn(value) {
			ans = append(ans, value)
		}
	}
	return ans
}

func (i Ints) Discard(fn convert) Ints {
	var ans Ints
	for _, value := range i {
		if !fn(value) {
			ans = append(ans, value)
		}
	}
	return ans
}

func (s Strings) Keep(fn stringConvert) Strings {
	var ans Strings
	for _, value := range s {
		if fn(value) {
			ans = append(ans, value)
		}
	}
	return ans
}

func (l Lists) Keep(fn listsConvert) Lists {
	var ans Lists
	for _, value := range l {
		if fn(value) {
			ans = append(ans, value)
		}
	}
	return ans
}
