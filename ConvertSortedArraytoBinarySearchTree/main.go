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

//TreeNode -
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	input := []int{-10, -3, 0, 5, 9}
	sortedArrayToBST(input)
}

func sortedArrayToBST(nums []int) *TreeNode {
	length := len(nums)
	if length == 1 {
		return &TreeNode{
			Val: nums[0],
		}
	}
	if length == 0 {
		return nil
	}

	return &TreeNode{
		Val:   nums[(length)/2],
		Left:  sortedArrayToBST(nums[0 : (length)/2]),
		Right: sortedArrayToBST(nums[(length)/2+1:]),
	}
}

// 注意点, left小于 node right 大于node
// 2和3的情况还是可以继续分解的,所以判断太多也会慢
