// 100. Same Tree
// Given two binary trees, write a function to check if they are the same or not.

// Two binary trees are considered the same if they are structurally identical and the nodes have the same value.

// Example 1:

// Input:     1         1
//           / \       / \
//          2   3     2   3

//         [1,2,3],   [1,2,3]

// Output: true
// Example 2:

// Input:     1         1
//           /           \
//          2             2

//         [1,2],     [1,null,2]

// Output: false
// Example 3:

// Input:     1         1
//           / \       / \
//          2   1     1   2

//         [1,2,1],   [1,1,2]

// Output: false
package main

import (
	"fmt"
	"reflect"
)

//TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	q := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isSameTree(p, q))
}

func isSameTree(p *TreeNode, q *TreeNode) bool {
	return reflect.DeepEqual(p, q)
}
