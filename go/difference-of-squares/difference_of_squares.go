package diffsquares

import (
	"math"
)

func SquareOfSums(r int) int {
	sum := 0
	for i := 0; i <= r; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func SumOfSquares(r int) int {
	squares := 0.0
	for i := 0; i <= r; i++ {
		squares += math.Pow(float64(i), 2)
	}
	return int(squares)
}

func Difference(r int) int {
	return SquareOfSums(r) - SumOfSquares(r)
}