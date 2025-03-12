package darts

func Score(x, y float64) int {
	position := x*x + y*y
	switch {
	case position <= 1:
		return 10
	case position <= 25:
		return 5
	case position <= 100:
		return 1
	default:
		return 0
	}
}
