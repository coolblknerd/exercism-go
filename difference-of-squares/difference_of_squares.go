package diffsquares

import (
	"math"
)

func SquareOfSum(num int) int {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(num int) int {
	sum := 0
	for i := 0; i <= num; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return sum
}

func Difference(num int) int {
	return SquareOfSum(num) - SumOfSquares(num)
}
