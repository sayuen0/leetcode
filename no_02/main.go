package main

import (
	"strconv"
	"strings"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 一つ一つを足し合わせる
	// できれば並列に計算したい
	a := add(l1) + add(l2)
	return intToReversedNode(a)
}

// 一個を足す
func add(l *ListNode) int {
	var sum int
	nl := l
	place := 1

	for {
		// 逆順から足していく
		sum += nl.Val * place
		if nl.Next == nil {
			break
		}
		place *= 10
		nl = nl.Next
	}
	return sum
}

// 逆順に格納したノードを返す
func intToReversedNode(v int) *ListNode {
	// intをバラして、一つの配列に収める
	arr := strings.Split(strconv.Itoa(v), "")
	root := &ListNode{}
	n := root
	for i := len(arr) - 1; i >= 0; i-- {
		n.Val, _ = strconv.Atoi(arr[i])
		if i > 0 {
			n.Next = &ListNode{}
			n = n.Next
		}
	}
	return root
}

// 配列を与えて、それを順番に持つノードを作成する
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
