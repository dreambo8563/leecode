package main

import (
	"fmt"
	"math"
)

/*
 * [7] Reverse Integer
 *
 * https://leetcode.com/problems/reverse-integer/description/
 *
 * algorithms
 * Easy (24.37%)
 * Total Accepted:    443.1K
 * Total Submissions: 1.8M
 * Testcase Example:  '123'
 *
 * Given a 32-bit signed integer, reverse digits of an integer.
 *
 * Example 1:
 *
 *
 * Input: 123
 * Output: 321
 *
 *
 * Example 2:
 *
 *
 * Input: -123
 * Output: -321
 *
 *
 * Example 3:
 *
 *
 * Input: 120
 * Output: 21
 *
 *
 * Note:
 * Assume we are dealing with an environment which could only store integers
 * within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of
 * this problem, assume that your function returns 0 when the reversed integer
 * overflows.
 *
 */

func main() {
	fmt.Println(reverse(123))
}
func reverse(x int) int {
	if x == 0 {
		return x
	}
	isNegative := x < 0

	num := x
	if isNegative {
		num = x * -1
	}
	result := 0

	for num > 0 {
		remainder := num % 10
		result *= 10
		result += remainder
		num /= 10
	}

	if isNegative {
		result *= -1
	}
	if result < math.MinInt32 || result > math.MaxInt32 {
		return 0
	}
	return result
}
