package main

import (
	"fmt"
)

/*
 * [28] Implement strStr()
 *
 * https://leetcode.com/problems/implement-strstr/description/
 *
 * algorithms
 * Easy (29.51%)
 * Total Accepted:    294.9K
 * Total Submissions: 998.4K
 * Testcase Example:  '"hello"\n"ll"'
 *
 * Implement strStr().
 *
 * Return the index of the first occurrence of needle in haystack, or -1 if
 * needle is not part of haystack.
 *
 * Example 1:
 *
 *
 * Input: haystack = "hello", needle = "ll"
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: haystack = "aaaaa", needle = "bba"
 * Output: -1
 *
 *
 * Clarification:
 *
 * What should we return when needle is an empty string? This is a great
 * question to ask during an interview.
 *
 * For the purpose of this problem, we will return 0 when needle is an empty
 * string. This is consistent to C's strstr() and Java's indexOf().
 *
 */
func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}

	c := needle[0]
	haystackLen := len(haystack)
	for i, s := range haystack {
		if byte(s) == c {
			remainLen := haystackLen - i
			// fmt.Println(i)
			// fmt.Println(remainLen)
			// 后面位数不够肯定不匹配
			if remainLen < needleLen {
				return -1
			}
			replacer := haystack[0:i] + needle[:] + haystack[i+needleLen:]
			if replacer == haystack {
				return i
			}
			// fmt.Printf("%v , %T \n", replacer, replacer)
		}
	}
	return -1
}

func main() {
	fmt.Println(strStr("hello", "ll"))
}
