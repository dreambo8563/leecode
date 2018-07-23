package main

import (
	"fmt"
	"sort"
)

/*
 * [21] Merge Two Sorted Lists
 *
 * https://leetcode.com/problems/merge-two-sorted-lists/description/
 *
 * algorithms
 * Easy (42.27%)
 * Total Accepted:    376.7K
 * Total Submissions: 890.2K
 * Testcase Example:  '[1,2,4]\n[1,3,4]'
 *
 * Merge two sorted linked lists and return it as a new list. The new list
 * should be made by splicing together the nodes of the first two lists.
 *
 * Example:
 *
 * Input: 1->2->4, 1->3->4
 * Output: 1->1->2->3->4->4
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode -
// 结果的 List 也是要排序的.....啊啊啊
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	vals := []int{}
	merge := &ListNode{}
	next := merge
	for l1 != nil {
		// fmt.Printf("%v", l1.Val)
		vals = append(vals, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		// fmt.Printf("%v", l2.Val)
		vals = append(vals, l2.Val)
		l2 = l2.Next
	}
	sort.Slice(vals, func(i, j int) bool { return vals[i] < vals[j] })
	for i, v := range vals {
		// fmt.Printf("%v", v)
		next.Val = v
		if i != len(vals)-1 {
			next.Next = &ListNode{}
			next = next.Next
		}
	}

	return merge
}

func main() {
	l1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 4,
			},
		},
	}
	l2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 3,
			Next: &ListNode{
				Val: 4,
			},
		},
	}

	merge := mergeTwoLists(l1, l2)
	for merge != nil {
		fmt.Printf("%v", merge.Val)
		// fmt.Println("1")
		merge = merge.Next
	}
}
