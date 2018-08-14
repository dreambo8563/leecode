// 118. Pascal's Triangle
// Given a non-negative integer numRows, generate the first numRows of Pascal's triangle.
// https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif
// In Pascal's triangle, each number is the sum of the two numbers directly above it.

// Example:

// Input: 5
// Output:
// [
//      [1],
//     [1,1],
//    [1,2,1],
//   [1,3,3,1],
//  [1,4,6,4,1]
// ]

package main

import (
	"fmt"
)

func main() {
	fmt.Println(generate(7))
}

func generate(numRows int) [][]int {
	init := [][]int{}
	if numRows == 0 {
		return init
	}

	var pre = make([]int, 1, 1)
	var tmp = make([]int, numRows, numRows)
	for index := 0; index < numRows; index++ {
		length := index + 1
		copy(tmp, pre)
		pre = make([]int, length, length)

		pre[0] = 1
		pre[length-1] = 1

		if index > 1 {
			fmt.Println(tmp)
			x := length % 2
			for c := 1; c < length/2+x; c++ {
				fmt.Printf("inner loop %v\n", c)
				pre[c] = tmp[c-1] + tmp[c]
				pre[len(pre)-1-c] = tmp[c-1] + tmp[c]
			}
		}
		init = append(init, pre)
	}
	return init
}
