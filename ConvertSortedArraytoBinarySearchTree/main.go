// 108. Convert Sorted Array to Binary Search Tree
// Given an array where elements are sorted in ascending order, convert it to a height balanced BST.

// For this problem, a height-balanced binary tree is defined as a binary tree in which the depth of the two subtrees of every node never differ by more than 1.

// Example:

// Given the sorted array: [-10,-3,0,5,9],

// One possible answer is: [0,-3,9,-10,null,5], which represents the following height balanced BST:

//       0
//      / \
//    -3   9
//    /   /
//  -10  5

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
	input := []int{-10, -3, 0, 5, 9}
	// tree := sortedArrayToBST(input)
	tree := sortedArrayToBST(input)
	printTree(tree)
}

// func sortedArrayToBST(nums []int) *TreeNode {
// 	return addNode(nums)
// }

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	fmt.Println(length)
	// if length == 3 {
	// 	return &TreeNode{
	// 		Val: nums[1],
	// 		Left: &TreeNode{
	// 			Val: nums[0],
	// 		},
	// 		Right: &TreeNode{
	// 			Val: nums[2],
	// 		},
	// 	}
	// }
	// if length == 2 {
	// 	return &TreeNode{
	// 		Val: nums[1],
	// 		Left: &TreeNode{
	// 			Val: nums[0],
	// 		},
	// 	}
	// }
	if length == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	if length == 0 {
		return nil
	}
	fmt.Printf("first part: %v \n", nums[0:(length-1)/2+1])
	fmt.Printf("second part: %v \n", nums[(length-1)/2+1:])

	// splitArr(nums[length/2:])
	return &TreeNode{
		Val:   nums[(length)/2],
		Left:  sortedArrayToBST(nums[0 : (length)/2]),
		Right: sortedArrayToBST(nums[(length)/2+1:]),
	}
}

func printTree(tree *TreeNode) {
	if tree != nil {
		fmt.Printf("tree %v \n", tree)
		if tree.Left != nil {
			fmt.Println("left node")
			printTree(tree.Left)
		}
		if tree.Right != nil {
			fmt.Println("right node")
			printTree(tree.Right)
		}
	}
}

// 注意点, left小于 node right 大于node
// 2和3的情况还是可以继续分解的,所以判断太多也会慢
