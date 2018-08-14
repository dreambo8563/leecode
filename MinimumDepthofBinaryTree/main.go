// 111. Minimum Depth of Binary Tree
// Given a binary tree, find its minimum depth.

// The minimum depth is the number of nodes along the shortest path from the root node down to the nearest leaf node.

// Note: A leaf is a node with no children.

// Example:

// Given binary tree [3,9,20,null,null,15,7],

//     3
//    / \
//   9  20
//     /  \
//    15   7
// return its minimum depth = 2.

package main

import (
	"fmt"
)

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
	fmt.Printf("depth: %d", minDepth(input))
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	level := 1
	nodeCheck([]*TreeNode{root}, &level)
	return level
}

func nodeCheck(arr []*TreeNode, level *int) (isEnd bool) {
	children := []*TreeNode{}
	for _, node := range arr {
		if node.Left == nil && node.Right == nil {
			return true
		}
		if node.Left != nil {
			children = append(children, node.Left)
		}
		if node.Right != nil {
			children = append(children, node.Right)
		}
	}
	*level++
	return nodeCheck(children, level)
}
