/*
   Largest prime factor

   The prime factors of 13195 are 5, 7, 13, and 29.

   What is the largest prime factor of the number 600851475143?
*/

package main

import "fmt"

func primeFactors(n int) (pfs []int) {
	// Get the number of 2s that divide n.
	for n%2 == 0 {
		pfs = append(pfs, 2)
		n = n / 2
	}
	// The number n must be odd at this point: skip one element.
	for i := 3; i*i <= n; i = i + 2 {
		// While i divides n, append i and divide n.
		for n%i == 0 {
			pfs = append(pfs, i)
			n = n / i
		}
	}
	// Handle cases when n is a prime number greater than 2.
	if n > 2 {
		pfs = append(pfs, n)
	}
	return
}

func largest(nums []int) (largest int) {
	for i := 0; i < len(nums); i++ {
		if nums[i] > largest {
			largest = nums[i]
		}
	}
	return
}

func main() {
	fmt.Printf("Largest is %d", largest(primeFactors(600851475143)))
}
