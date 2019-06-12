package main

import (
	"fmt"
)

/*
 * [35] Search Insert Position
 *
 * https://leetcode.com/problems/search-insert-position/description/
 *
 * algorithms
 * Easy (40.03%)
 * Total Accepted:    278.4K
 * Total Submissions: 695.5K
 * Testcase Example:  '[1,3,5,6]\n5'
 *
 * Given a sorted array and a target value, return the index if the target is
 * found. If not, return the index where it would be if it were inserted in
 * order.
 *
 * You may assume no duplicates in the array.
 *
 * Example 1:
 *
 *
 * Input: [1,3,5,6], 5
 * Output: 2
 *
 *
 * Example 2:
 *
 *
 * Input: [1,3,5,6], 2
 * Output: 1
 *
 *
 * Example 3:
 *
 *
 * Input: [1,3,5,6], 7
 * Output: 4
 *
 *
 * Example 4:
 *
 *
 * Input: [1,3,5,6], 0
 * Output: 0
 *
 *
 */
func searchInsert(nums []int, target int) int {
	if len(nums) == 0 || target < nums[0] {
		return 0
	}
	start := 0
	end := len(nums) - 1
	if target > nums[end] {
		return len(nums)
	}

	mid := 0

	for {
		diff := end - start
		fmt.Printf("start diff: %d, start: %d, end: %d \n", diff, start, end)
		if diff == 0 {
			return start
		}
		mid = (diff >> 1) + start
		if target == nums[mid] {
			for mid >= 0 && target == nums[mid] && mid >= start {
				fmt.Printf("inside loop, mid: %d \n", mid)
				mid--
			}
			return mid + 1
		}
		if target > nums[mid] {
			fmt.Printf("target > nums[mid], mid: %d \n", mid)
			if start == mid {
				start++
			} else {
				start = mid
			}
			continue
		} else {
			fmt.Printf("NOT target > nums[mid], mid: %d \n", mid)
			end = mid
			continue
		}
	}
}

func main() {
	fmt.Println(searchInsert([]int{1, 3}, 1))
}
