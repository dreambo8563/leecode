/*
 * [53] Maximum Subarray
 *
 * https://leetcode.com/problems/maximum-subarray/description/
 *
 * algorithms
 * Easy (40.80%)
 * Total Accepted:    340.1K
 * Total Submissions: 832.8K
 * Testcase Example:  '[-2,1,-3,4,-1,2,1,-5,4]'
 *
 * Given an integer array nums, find the contiguous subarray (containing at
 * least one number) which has the largest sum and return its sum.
 *
 * Example:
 *
 *
 * Input: [-2,1,-3,4,-1,2,1,-5,4],
 * Output: 6
 * Explanation: [4,-1,2,1] has the largest sum = 6.
 *
 *
 * Follow up:
 *
 * If you have figured out the O(n) solution, try coding another solution using
 * the divide and conquer approach, which is more subtle.
 *
 */

package main

import (
	"fmt"
)

func main() {
	input := []int{2, 0, 3, -2}
	// input := []int{-1, -3, -9, -2, 8}
	fmt.Println(maxSubArray(input))
}

func maxSubArray(nums []int) int {
	content := filterContent(nums)
	if len(content) == 0 {
		maxValue := nums[0]
		for index := 0; index < len(nums)-1; index++ {
			maxValue = max(maxValue, nums[index+1])
		}
		fmt.Println(maxValue)
		return maxValue
	}
	fmt.Printf("filter content: %v \n", content)
	max := innerSum(content)
	return max
}

func innerSum(nums []int) int {
	maxValue := nums[0]
	tmp := 0
	for index := 1; index < len(nums); index++ {
		if nums[index] >= 0 {
			if nums[index]+maxValue+tmp <= maxValue {
				tmp += nums[index]
			} else {
				maxValue = nums[index] + maxValue + tmp
				tmp = 0
			}

			fmt.Printf("nums[index] > 0: %v  maxValue:%v\n", nums[index], maxValue)
			continue
		}
		if nums[index] < 0 && nums[index]+maxValue+tmp > 0 {
			tmp += nums[index]
			fmt.Printf("nums[index]:%v, tmp保存后: %v,   maxValue:%v \n", nums[index], tmp, maxValue)
			continue
		}
		if nums[index] < 0 && nums[index]+maxValue+tmp == 0 {
			tmp += nums[index]
			fmt.Printf("tmp保存后: %v,   maxValue:%v \n", tmp, maxValue)
			continue
		}
		tmp = 0
		maxValue = max(maxSubArray(nums[index+1:]), maxValue)
		fmt.Printf("inner loop max %v \n", maxValue)
		break
	}
	fmt.Printf("阶段性max %v \n", maxValue)
	return maxValue
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func filterContent(inputs []int) []int {
	var begin, end int
	// fmt.Println(begin)
	for index, value := range inputs {
		if value > 0 {
			begin = index
			break
		}
	}
	inputsLen := len(inputs)
	// fmt.Println(begin)
	// fmt.Println(inputsLen)
	if inputs[begin] <= 0 || begin == inputsLen-1 {
		return []int{}
	}

	for index := range inputs {
		if inputs[inputsLen-1-index] > 0 {
			end = inputsLen - 1 - index
			break
		}
	}

	if begin == end {
		return []int{}
	}

	return inputs[begin : end+1]
}
