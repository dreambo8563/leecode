// 107. Binary Tree Level Order Traversal II
// Given a binary tree, return the bottom-up level order traversal of its nodes' values. (ie, from left to right, level by level from leaf to root).

// For example:
// Given binary tree [3,9,20,null,null,15,7],
//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its bottom-up level order traversal as:
// [
//   [15,7],
//   [9,20],
//   [3]
// ]

package main

import "fmt"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

//TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	input := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	fmt.Println(levelOrderBottom(input))
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	result := [][]int{[]int{root.Val}}
	children := []*TreeNode{root}
	ended := true
	for {
		ended, children = hasChild(children, &result)
		if ended {
			break
		}
	}
	// for left, right := 0, len(result)-1; left < right; left, right = left+1, right-1 {
	// 	result[left], result[right] = result[right], result[left]
	// }
	return result
}

func hasChild(arr []*TreeNode, result *[][]int) (ended bool, children []*TreeNode) {
	ended = true
	if len(arr) == 0 {
		return true, []*TreeNode{}
	}
	innerS := []int{}
	for _, node := range arr {

		if node.Left != nil {
			children = append(children, node.Left)
			innerS = append(innerS, node.Left.Val)
		}
		if node.Right != nil {
			children = append(children, node.Right)
			innerS = append(innerS, node.Right.Val)
		}

	}
	if len(innerS) > 0 {
		*result = append([][]int{innerS}, (*result)...)
	}
	if len(children) > 0 {
		ended = false
	}
	return
}
