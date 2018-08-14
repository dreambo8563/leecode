// 110. Balanced Binary Tree

// Given a binary tree, determine if it is height-balanced.

// For this problem, a height-balanced binary tree is defined as:

// a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

// Example 1:

// Given the following tree [3,9,20,null,null,15,7]:

//     3
//    / \
//   9  20
//     /  \
//    15   7
// Return true.

// Example 2:

// Given the following tree [1,2,2,3,3,null,null,4,4]:

//        1
//       / \
//      2   2
//     / \
//    3   3
//   / \
//  4   4
// Return false.

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
	// input := &TreeNode{
	// 	Val: 3,
	// 	Left: &TreeNode{
	// 		Val: 9,
	// 		Left: &TreeNode{
	// 			Val: 15,
	// 			Left: &TreeNode{
	// 				Val: 66,
	// 			},
	// 			Right: &TreeNode{
	// 				Val: 7,
	// 			},
	// 		},
	// 		Right: &TreeNode{
	// 			Val: 7,
	// 		},
	// 	},
	// 	Right: &TreeNode{
	// 		Val: 20,
	// 	},
	// }
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
	fmt.Printf("The Binary Tree isBalanced: %t\n", isBalanced(input))
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	_, b := getDepth(root)
	return b
}

func getDepth(root *TreeNode) (depth int, isBalanced bool) {
	isBalanced = true
	if root == nil {
		depth = 0
		return
	}
	depth = 1
	lDepth, lb := getDepth(root.Left)
	rDepth, rb := getDepth(root.Right)

	if !lb || !rb {
		fmt.Printf("not balance l or r: l-%v r-%v \n", root.Left, root.Right)
		isBalanced = false
		return
	}

	if (max(lDepth, rDepth) - min(lDepth, rDepth)) > 1 {
		fmt.Printf("diff to much: %v %v \n", root.Left, root.Right)
		isBalanced = false
		return
	}
	depth += max(lDepth, rDepth)
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 多次提交失败发现问题是对平衡树的定义不清楚
// 只要保持两个子节点的最大深度差距不超过1 就行
// 不要拿子节点里的最大深度和最小深度比,而是要让左右子节点的最大深度来比
