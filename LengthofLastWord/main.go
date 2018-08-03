/*
 * [58] Length of Last Word
 *
 * https://leetcode.com/problems/length-of-last-word/description/
 *
 * algorithms
 * Easy (32.07%)
 * Total Accepted:    204.6K
 * Total Submissions: 637.8K
 * Testcase Example:  '"Hello World"'
 *
 * Given a string s consists of upper/lower-case alphabets and empty space
 * characters ' ', return the length of last word in the string.
 *
 * If the last word does not exist, return 0.
 *
 * Note: A word is defined as a character sequence consists of non-space
 * characters only.
 *
 * Example:
 *
 * Input: "Hello World"
 * Output: 5
 *
 *
 */

package main

import (
	"fmt"
)

func main() {
	input := "a "
	result := lengthOfLastWord(input)
	fmt.Printf("the length of last word is : %d", result)
}
func lengthOfLastWord(s string) int {
	strLen := len(s)
	fmt.Println(strLen)
	start := strLen
	end := strLen
	for index := strLen - 1; index >= 0; index-- {
		if s[index] != 32 && end == strLen {
			end = index
			fmt.Printf("find end: %s\n", string(s[end]))
			continue
		}
		if s[index] == 32 && end != strLen {
			start = index + 1
			fmt.Printf("find start: %s\n", string(s[start]))
			break
		}
	}
	if start == end && start == strLen && end == strLen {
		return 0
	}
	if start == strLen && end != strLen {
		start = 0
	}
	if start == end {
		return 1
	}
	fmt.Printf("start: %d, end: %d \n", start, end)
	return end - start + 1
}
