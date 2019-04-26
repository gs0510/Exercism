package diffsquares

// SquareOfSum calculates the square of sum of first N natural numbers.
func SquareOfSum(n int) int {
	return ((n * (n + 1)) / 2) * ((n * (n + 1)) / 2)
}

// SumOfSquares calculates the sum of squares of first N natural numbers.
func SumOfSquares(n int) int {
	return (n * (n + 1) * (2*n + 1)) / 6
}

// Difference calculates the difference between square of sum and sum of sqaures of first N natural numbers.
func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
