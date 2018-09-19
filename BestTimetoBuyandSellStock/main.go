/*
 * [121] Best Time to Buy and Sell Stock
 *
 * https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/
 *
 * algorithms
 * Easy (44.30%)
 * Total Accepted:    346.8K
 * Total Submissions: 782.9K
 * Testcase Example:  '[7,1,5,3,6,4]'
 *
 * Say you have an array for which the ith element is the price of a given
 * stock on day i.
 *
 * If you were only permitted to complete at most one transaction (i.e., buy
 * one and sell one share of the stock), design an algorithm to find the
 * maximum profit.
 *
 * Note that you cannot sell a stock before you buy one.
 *
 * Example 1:
 *
 *
 * Input: [7,1,5,3,6,4]
 * Output: 5
 * Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit
 * = 6-1 = 5.
 * Not 7-1 = 6, as selling price needs to be larger than buying price.
 *
 *
 * Example 2:
 *
 *
 * Input: [7,6,4,3,1]
 * Output: 0
 * Explanation: In this case, no transaction is done, i.e. max profit = 0.
 *
 *
 */

package main

import (
	"fmt"
)

func main() {
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit(input))
}

func maxProfit(prices []int) int {
	length := len(prices)
	if length < 2 {
		return 0
	}
	var min, max = prices[0], prices[1]

	if prices[0] > prices[1] {
		max = prices[1]
		min = prices[1]
	}

	remain := prices[2:]
	for k, v := range remain {
		fmt.Printf("index: %v, value: %v \n", k, v)
		if v < min {
			// min = v
			fmt.Printf("diff %v, cur %v \n", max-min, maxProfit(remain[k:]))
			return maxInt(max-min, maxProfit(remain[k:]))
		}
		if v > max {
			max = v
			continue
		}
	}
	fmt.Printf("min: %v, max: %v \n", min, max)
	// result := 0
	// if maxIndex > minIndex {
	// 	result =
	// }
	return max - min
}

func maxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 别人更好的解法

// func maxProfit(prices []int) int {
// 	max_profit := 0
// 	size := len(prices)
// 	maxPrice := make([]int, size)
// 	max_till_now := 0
// 从后到前, 每个元素都是他之后所有元素的最大值
// 	for i := size - 1; i >= 0; i-- {
// 		maxPrice[i] = max_till_now
// 		if max_till_now < prices[i] {
// 			max_till_now = prices[i]
// 		}
// 	}

// 每一个元素都和后面的最大值比较,拿到差的最大值
// 	for i := 0; i < size; i++ {
// 		if (maxPrice[i] - prices[i]) > max_profit {
// 			max_profit = maxPrice[i] - prices[i]
// 		}
// 	}
// 	return max_profit
// }
