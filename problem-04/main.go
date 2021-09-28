/*
   Largest palindrome product

   A palindromic number reads the same both ways.

   The largest palindrom made from the product of two 2-digit numbers is 9009 = 91 x 99.

   Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import "fmt"

func reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

func isPalindrome(s string) bool {
	return s == reverse(s)
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
	nums := make([]int, 0)
	for i := 999; i > 100; i-- {
		for j := 999; i > 100; j-- {
			m := j * i
			if isPalindrome(string(m)) {
				nums = append(nums, m)
			}
		}
	}
	fmt.Printf("Largest is %d", largest(nums))
}
