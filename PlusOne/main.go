/*
 * [66] Plus One
 *
 * https://leetcode.com/problems/plus-one/description/
 *
 * algorithms
 * Easy (39.89%)
 * Total Accepted:    270.9K
 * Total Submissions: 678.6K
 * Testcase Example:  '[1,2,3]'
 *
 * Given a non-empty array of digits representing a non-negative integer, plus
 * one to the integer.
 *
 * The digits are stored such that the most significant digit is at the head of
 * the list, and each element in the array contain a single digit.
 *
 * You may assume the integer does not contain any leading zero, except the
 * number 0 itself.
 *
 * Example 1:
 *
 *
 * Input: [1,2,3]
 * Output: [1,2,4]
 * Explanation: The array represents the integer 123.
 *
 *
 * Example 2:
 *
 *
 * Input: [4,3,2,1]
 * Output: [4,3,2,2]
 * Explanation: The array represents the integer 4321.
 *
 *
 */

package main

import (
	"fmt"
)

func main() {
	input := []int{9}
	result := plusOne(input)
	fmt.Println(result)

}

func plusOne(digits []int) []int {
	inputLen := len(digits)
	promotion := 1
	for index := inputLen - 1; index >= 0; index-- {
		if digits[index] == 9 && promotion == 1 {
			if index == 0 {
				tmp := make([]int, inputLen)
				digits = append([]int{1}, tmp...)
				break
			}
			digits[index] = 0
			promotion = 1
		} else {
			digits[index]++
			break
		}
	}
	return digits
}
