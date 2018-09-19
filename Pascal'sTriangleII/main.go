// 119. Pascal's Triangle II
// Given a non-negative index k where k ≤ 33, return the kth index row of the Pascal's triangle.

// Note that the row index starts from 0.

// In Pascal's triangle, each number is the sum of the two numbers directly above it.

// Example:

// Input: 3
// Output: [1,3,3,1]

package main

import "fmt"

func main() {
	fmt.Println(getRow(0))
}
func getRow(rowIndex int) []int {
	// if rowIndex == 0 {
	// 	return []int{}
	// }
	triangles := generate(rowIndex + 1)
	// fmt.Println(triangles)
	return triangles[len(triangles)-1]
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
			// fmt.Println(tmp)
			x := length % 2
			for c := 1; c < length/2+x; c++ {
				// fmt.Printf("inner loop %v\n", c)
				pre[c] = tmp[c-1] + tmp[c]
				pre[len(pre)-1-c] = tmp[c-1] + tmp[c]
			}
		}
		init = append(init, pre)
	}
	return init
}
