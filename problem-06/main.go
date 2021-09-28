/*
   Sum square difference

   The sum of the squares of the first ten natural numbers is,

   1^2 + 2^2 + ... + 10^2 = 385

   The square of the sum of the first ten natural numbers is,

   (1 + 2 + ... + 10)^2 = 55^2 = 3025

   Hence the difference between the sum of the squares of the first ten
   natural numbers and the square of the sum is 3025 - 385 = 2640.

   Find the difference between the sum of the squares of the first one hundred
   natural numbers and the square of the sum.
*/
package main

import "fmt"

func sumOfSquares(n int) (sum int) {
	for i := 0; i <= n; i++ {
		sum += i * i
	}
	return
}

func squareOfSum(n int) (sq int) {
	sum := 0
	for i := 0; i <= n; i++ {
		sum += i
	}
	sq = sum * sum
	return
}

func main() {
	// sum, sq :=
	// 	sumOfSquares(10),
	// 	squareOfSum(10)

	sum, sq :=
		sumOfSquares(100),
		squareOfSum(100)

	diff := sq - sum
	fmt.Println(sum, sq, diff)
}
