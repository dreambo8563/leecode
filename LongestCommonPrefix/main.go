package main

import (
	"fmt"
)

/*
 * [14] Longest Common Prefix
 *
 * https://leetcode.com/problems/longest-common-prefix/description/
 *
 * algorithms
 * Easy (31.76%)
 * Total Accepted:    298.8K
 * Total Submissions: 940.4K
 * Testcase Example:  '["flower","flow","flight"]'
 *
 * Write a function to find the longest common prefix string amongst an array
 * of strings.
 *
 * If there is no common prefix, return an empty string "".
 *
 * Example 1:
 *
 *
 * Input: ["flower","flow","flight"]
 * Output: "fl"
 *
 *
 * Example 2:
 *
 *
 * Input: ["dog","racecar","car"]
 * Output: ""
 * Explanation: There is no common prefix among the input strings.
 *
 *
 * Note:
 *
 * All given inputs are in lowercase letters a-z.
 *
 */
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	min := len(strs[0])
	pre := strs[0]
	for index, str := range strs {
		if index == 0 {
			continue
		}
		if len(str) < min {
			// update min lenght
			min = len(str)
		}
		// fmt.Println(string(pre[i]))
		fmt.Println(min)
		for i := 0; i < min; i++ {

			if pre[i] != str[i] {
				min = i
				fmt.Println("not equal")
				fmt.Println(min)
				pre = pre[:i]
				break
			}
		}
		pre = pre[:min]
		fmt.Println(pre + "after loop")
	}
	return pre
}

func main() {
	str := []string{}
	fmt.Println(longestCommonPrefix(str))
}
