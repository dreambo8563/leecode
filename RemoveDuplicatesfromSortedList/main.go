//83. Remove Duplicates from Sorted List

// Given a sorted linked list, delete all duplicates such that each element appear only once.

// Example 1:

// Input: 1->1->2
// Output: 1->2
// Example 2:

// Input: 1->1->2->3->3
// Output: 1->2->3
package main

import (
	"fmt"
)

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode -
type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{
		Val: 1, Next: &ListNode{
			Val: 2, Next: nil,
		},
	},
	}
	deleteDuplicates(head)
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	currentNode := head
	var li = &ListNode{Val: currentNode.Val, Next: currentNode.Next}
	nextNode := li

	for currentNode != nil {
		fmt.Printf("current Val %d \n", currentNode.Val)
		if currentNode.Next != nil {
			next := currentNode.Next.Val
			if next != currentNode.Val {
				fmt.Printf("next Val %d \n", currentNode.Next.Val)
				nextNode.Next = currentNode.Next
				nextNode = nextNode.Next
				// continue
				// nextNode = currentNode.Next
			}
			currentNode = currentNode.Next
		} else {
			nextNode.Next = nil
			break
		}
		// printList(li)

	}
	// printList(li)
	fmt.Printf("nextNode val, %d \n", nextNode)
	fmt.Printf("currentNode val, %d \n", currentNode)
	// li.Next = nil
	return li
}
func printList(li *ListNode) {
	for li != nil {
		fmt.Println(li.Val)
		li = li.Next
	}
}
