/*
 * [67] Add Binary
 *
 * https://leetcode.com/problems/add-binary/description/
 *
 * algorithms
 * Easy (34.99%)
 * Total Accepted:    220.1K
 * Total Submissions: 625.2K
 * Testcase Example:  '"11"\n"1"'
 *
 * Given two binary strings, return their sum (also a binary string).
 *
 * The input strings are both non-empty and contains only characters 1 or 0.
 *
 * Example 1:
 *
 *
 * Input: a = "11", b = "1"
 * Output: "100"
 *
 * Example 2:
 *
 *
 * Input: a = "1010", b = "1011"
 * Output: "10101"
 *
 */

package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("reuslt: %s", addBinary("11", "1"))
}

func addBinary(a string, b string) string {
	x, y := align(a, b)
	const one = 49
	// const zero = 48
	fmt.Printf("x: %s, y: %s \n", x, y)
	promotion := 0
	result := ""
	for index := len(x) - 1; index >= 0; index-- {
		if x[index] == one && y[index] == one {

			result = strconv.Itoa(0+promotion) + result

			promotion = 1
			fmt.Printf(" && a: %d ,b: %d , result: %s,promotion:%d \n", x[index], y[index], result, promotion)
			continue
		}
		if x[index] == one || y[index] == one {
			if promotion == 1 {
				promotion = 1
				result = "0" + result
				fmt.Printf("|| with P -  a: %d ,b: %d , result: %s, promotion:%d \n", x[index], y[index], result, promotion)
				continue
			}
			promotion = 0
			result = "1" + result
			fmt.Printf("|| wt P - a: %d ,b: %d , result: %s, promotion:%d \n", x[index], y[index], result, promotion)
			continue
		}

		result = strconv.Itoa(0+promotion) + result
		fmt.Printf("a: %d ,b: %d , result: %s,promotion:%d  \n", x[index], y[index], result, promotion)
		promotion = 0
	}
	if promotion == 1 {
		result = "1" + result
	}
	return result
}

func align(a, b string) (x, y string) {
	aLen := len(a)
	bLen := len(b)
	if aLen == bLen {
		return a, b
	}
	diff := math.Abs(float64(aLen - bLen))
	if aLen > bLen {
		return a, strings.Repeat("0", int(diff)) + b
	}
	return strings.Repeat("0", int(diff)) + a, b
}
