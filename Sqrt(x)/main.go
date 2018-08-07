/*
 * [69] Sqrt(x)
 *
 * https://leetcode.com/problems/sqrtx/description/
 *
 * algorithms
 * Easy (29.27%)
 * Total Accepted:    255.7K
 * Total Submissions: 871.7K
 * Testcase Example:  '4'
 *
 * Implement int sqrt(int x).
 *
 * Compute and return the square root of x, where x is guaranteed to be a
 * non-negative integer.
 *
 * Since the return type is an integer, the decimal digits are truncated and
 * only the integer part of the result is returned.
 *
 * Example 1:
 *
 *
 * Input: 4
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: 8
 * Output: 2
 * Explanation: The square root of 8 is 2.82842..., and since
 * the decimal part is truncated, 2 is returned.
 *
 *
 */

package main

import (
	"fmt"
)

func main() {
	fmt.Printf("sqrt(int x): %d", mySqrt(8))
}

//Newton's method
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	xF := float64(x)

	return bingo(xF/2, xF)
}

func bingo(m float64, init float64) int {
	result := (m + init/m) / float64(2.0)
	fmt.Printf("m: %f , result: %f \n", m, result)
	if int(m) == int(result) {
		return int(m)
	}
	return bingo(result, init)
}
