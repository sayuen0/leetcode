package main

import (
	"bytes"
	"strconv"
)

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

func (l *ListNode) String() string {
	out := bytes.Buffer{}
	out.WriteString("[")
	current := l
	for current != nil {
		out.WriteString(strconv.Itoa(current.Val))
		if current.Next != nil {
			out.WriteString(",")
		}
		current = current.Next
	}
	out.WriteString("]")
	return out.String()
}
