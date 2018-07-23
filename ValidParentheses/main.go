package main

import (
	"fmt"
)

/*
 * [20] Valid Parentheses
 *
 * https://leetcode.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (34.32%)
 * Total Accepted:    365.3K
 * Total Submissions: 1.1M
 * Testcase Example:  '"()"'
 *
 * Given a string containing just the characters '(', ')', '{', '}', '[' and
 * ']', determine if the input string is valid.
 *
 * An input string is valid if:
 *
 *
 * Open brackets must be closed by the same type of brackets.
 * Open brackets must be closed in the correct order.
 *
 *
 * Note that an empty string is also considered valid.
 *
 * Example 1:
 *
 *
 * Input: "()"
 * Output: true
 *
 *
 * Example 2:
 *
 *
 * Input: "()[]{}"
 * Output: true
 *
 *
 * Example 3:
 *
 *
 * Input: "(]"
 * Output: false
 *
 *
 * Example 4:
 *
 *
 * Input: "([)]"
 * Output: false
 *
 *
 * Example 5:
 *
 *
 * Input: "{[]}"
 * Output: true
 *
 *
 */
var brackets = map[rune]int{
	'(': 1,
	')': -1,
	'[': 2,
	']': -2,
	'{': 3,
	'}': -3,
}

// var brackets = [6]int{'(', '[', '{', ')', ']', '}'}

func isValid(s string) bool {

	if s == "" {
		return true
	}
	if len(s)%2 != 0 {
		return false
	}
	stack := []int{}
	for _, r := range s {
		if brackets[r] > 0 {
			stack = append(stack, brackets[r])
			continue
		} else {
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			// fmt.Println(stack)
			if brackets[r] != 0-last {
				return false
			}
			stack = stack[:len(stack)-1]
			// fmt.Println(stack)
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}

func main() {
	result := isValid("[")
	fmt.Println(result)
}
