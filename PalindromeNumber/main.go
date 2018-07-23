package main

import "fmt"

/*
 * [9] Palindrome Number
 *
 * https://leetcode.com/problems/palindrome-number/description/
 *
 * algorithms
 * Easy (37.27%)
 * Total Accepted:    361.9K
 * Total Submissions: 969.8K
 * Testcase Example:  '121'
 *
 * Determine whether an integer is a palindrome. An integer is a palindrome
 * when it reads the same backward as forward.
 *
 * Example 1:
 *
 *
 * Input: 121
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: -121
 * Output: false
 * Explanation: From left to right, it reads -121. From right to left, it
 * becomes 121-. Therefore it is not a palindrome.
 *
 *
 * Example 3:
 *
 *
 * Input: 10
 * Output: false
 * Explanation: Reads 01 from right to left. Therefore it is not a
 * palindrome.
 *
 *
 * Follow up:
 *
 * Coud you solve it without converting the integer to a string?
 *
 */
func main() {
	fmt.Println(isPalindrome(10))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}
	length := 0
	arr := []int{}
	for x > 0 {
		length++
		arr = append(arr, x%10)
		x /= 10
	}

	mid := length / 2
	// fmt.Println(arr[0:mid])
	// fmt.Println(arr[len(arr)-mid:])
	compose := append(arr[0:mid], arr[len(arr)-mid:]...)
	for k, v := range compose {

		if v != compose[len(compose)-k-1] {
			// fmt.Printf("%v + %v  \n", compose[len(compose)-k-1], v)
			return false
		}
	}
	return true
}
