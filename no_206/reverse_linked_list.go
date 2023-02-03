package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	return helper(head, nil)
}

func helper(current, prev *ListNode) *ListNode {
	if current == nil {
		return prev
	}

	next := current.Next
	current.Next = prev
	return helper(next, current)
}
