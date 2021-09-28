/*
   10001st prime

   By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the
   6th place prime is 13.

   What is the 10001st prime number?
*/
package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	for i := 2; i <= int(math.Floor(float64(n)/2)); i++ {
		if n%i == 0 {
			return false
		}
	}
	return n > 1
}

func primeAtIndex(n int) (nums []int) {
	count := 0
	for i := 0; count < n; i++ {
		if isPrime(i) {
			count++
			if count == n {
				nums = append(nums, i)
			}
		}
	}
	return
}

func main() {
	// fmt.Println(primeAtIndex(6))
	fmt.Println(primeAtIndex(10001))
}
