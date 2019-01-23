package main

import (
	"fmt"
	"sort"
)

/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 *
 * https://leetcode.com/problems/single-number/description/
 *
 * algorithms
 * Easy (58.39%)
 * Total Accepted:    398.5K
 * Total Submissions: 682.3K
 * Testcase Example:  '[2,2,1]'
 *
 * Given a non-empty array of integers, every element appears twice except for
 * one. Find that single one.
 *
 * Note:
 *
 * Your algorithm should have a linear runtime complexity. Could you implement
 * it without using extra memory?
 *
 * Example 1:
 *
 *
 * Input: [2,2,1]
 * Output: 1
 *
 *
 * Example 2:
 *
 *
 * Input: [4,1,2,1,2]
 * Output: 4
 *
 *
 */

func main() {
	input := []int{2, 2, 1}[:]
	fmt.Println(singleNumber(input))
}

func singleNumber(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	sort.Slice(nums, func(a, b int) bool {

		return nums[a] < nums[b]
	})

	for index := 0; index+1 < length; {
		fmt.Println(nums[index], nums[index+1])
		if nums[index] != nums[index+1] {
			return nums[index]
		}
		index = index + 2
	}
	return nums[len(nums)-1]
}
