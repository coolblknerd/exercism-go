package strain

// Elements is a public type
type Elements []interface{}

// Keep is a public function
func (values Elements) Keep(f func(interface{}) bool) Elements {
	var strain Elements
	for _, elem := range values {
		if f(elem) {

			strain = append(strain, elem)
		}
	}
	return strain
}

// // Discard is a public function
// func (nums Ints) Discard(f func(int) bool) Ints {
// 	var strain Ints
// 	for _, num := range nums {
// 		if !f(num) {
// 			strain = append(strain, num)
// 		}
// 	}
// 	return strain
// }
