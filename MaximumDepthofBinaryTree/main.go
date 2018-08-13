// 104. Maximum Depth of Binary Tree
// Given a binary tree, find its maximum depth.

// The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

// Note: A leaf is a node with no children.

// Example:

// Given binary tree [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its depth = 3.

package main

import (
	"fmt"
)

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
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 4,
			},
		},
	}

	fmt.Println(maxDepth(input))
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 1
	children := []*TreeNode{root}
	ended := true
	for {
		ended, children = hasChild(children, &level)
		if ended {
			break
		}
	}
	return level
}

func hasChild(arr []*TreeNode, level *int) (ended bool, children []*TreeNode) {
	ended = true
	if len(arr) == 0 {
		return true, []*TreeNode{}
	}
	for _, node := range arr {
		if node.Left != nil {
			children = append(children, node.Left)
		}
		if node.Right != nil {
			children = append(children, node.Right)
		}
	}

	if len(children) > 0 {
		*level++
		ended = false
	}
	return
}
