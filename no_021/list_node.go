package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(arr ...int) *ListNode {
	n := &ListNode{}
	root := n
	for i, v := range arr {
		root.Val = v
		if i != len(arr)-1 {
			root.Next = &ListNode{}
			root = root.Next
		}
	}
	return n
}
