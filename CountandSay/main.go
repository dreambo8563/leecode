package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
 * [38] Count and Say
 *
 * https://leetcode.com/problems/count-and-say/description/
 *
 * algorithms
 * Easy (37.38%)
 * Total Accepted:    207.3K
 * Total Submissions: 553.8K
 * Testcase Example:  '1'
 *
 * The count-and-say sequence is the sequence of integers with the first five
 * terms as following:
 *
 * 1.     1
 * 2.     11
 * 3.     21
 * 4.     1211
 * 5.     111221
 *
 *
 *
 * 1 is read off as "one 1" or 11.
 * 11 is read off as "two 1s" or 21.
 * 21 is read off as "one 2, then one 1" or 1211.
 *
 *
 *
 * Given an integer n, generate the nth term of the count-and-say sequence.
 *
 *
 *
 * Note: Each term of the sequence of integers will be represented as a
 * string.
 *
 *
 * Example 1:
 *
 * Input: 1
 * Output: "1"
 *
 *
 *
 * Example 2:
 *
 * Input: 4
 * Output: "1211"
 *
 *
 */
func countAndSay(n int) string {
	countStr := "1"
	for index := 1; index < n; index++ {
		countStr = say(countStr)
	}
	return countStr
}

func say(input string) string {
	raw := []rune(input)
	sound := ""

	for len(raw) != 0 {
		result, rest := splitBegin(raw)
		fmt.Printf("result: %s, rest: %v \n", result, rest)
		raw = rest
		sound += result
	}
	return sound
}

func splitBegin(input []rune) (string, []rune) {
	if len(input) == 0 {
		log.Fatalln("input is empty")
	}
	first := input[0]
	for i, c := range input {
		if c != first {
			return strconv.Itoa(i) + string(first), input[i:]
		}
	}
	fmt.Printf("len(input) + first: %v \n", strconv.Itoa(len(input))+string(first))
	return strconv.Itoa(len(input)) + string(first), []rune{}
}

func main() {
	str := countAndSay(10)
	fmt.Println(str)
}
