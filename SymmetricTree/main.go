// 101. Symmetric Tree
// Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

// For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

//     1
//    / \
//   2   2
//  / \ / \
// 3  4 4  3
// But the following [1,2,2,null,3,null,3] is not:
//     1
//    / \
//   2   2
//    \   \
//    3    3
// Note:
// Bonus points if you could solve it both recursively and iteratively.

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
		Val: 1,
		Left: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		}, Right: &TreeNode{
			Val:  2,
			Left: nil,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(isSymmetric(input))
}

func isSymmetric(root *TreeNode) bool {
	// fmt.Println(*root)
	if root == nil {
		return true
	}
	vl, fl := composeList([]*TreeNode{root})

	// level := 1
	for {
		fmt.Printf("len(fl): %d \n", len(fl))
		if len(fl)%2 != 0 {
			return false
		}
		if len(fl) == 0 {
			break
		}

		fmt.Printf("new fl: %x \n", fl)
		if isMirror(vl) {
			fmt.Printf("vl: %v \n", vl)
			vl, fl = composeList(fl)
			continue
		}
		return false
	}
	// fmt.Printf("%x", tmpList)
	return true
}

func composeList(tl []*TreeNode) (l []int, fl []*TreeNode) {
	length := len(tl)

	for index := 0; index < length; index++ {
		if tl[index] == nil && tl[length-1-index] != nil {
			fmt.Println("nil")
			return []int{1}, []*TreeNode{&TreeNode{Val: 9}}
		}
		if tl[index] != nil {
			l = append(l, tl[index].Val)
			// fmt.Println(l)
			// if tl[index].Left != nil {

			// }
			fl = append(fl, tl[index].Left)

			// if tl[index].Right != nil {

			// }
			fl = append(fl, tl[index].Right)
		}

	}
	return
}

func isMirror(input []int) bool {
	length := len(input)
	for index := 0; index < length-1; index++ {
		if input[index] == input[length-1-index] {
			continue
		}
		return false
	}

	return true
}

// 人家的递归的好啊
// func isSym(n1, n2 *TreeNode) bool {
//     if n1 == nil && n2 == nil {
//         return true
//     }
//     if n1 != nil && n2 != nil && n1.Val == n2.Val {
//         return isSym(n1.Left, n2.Right) && isSym(n1.Right, n2.Left)
//     }
//     return false
// }

// func isSymmetric(root *TreeNode) bool {
//     return isSym(root, root)

// }
