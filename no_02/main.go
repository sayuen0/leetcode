package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// 右辺と左辺
	out := &ListNode{}
	current := out
	mod := 0
	leftVal, rightVal, sum := 0, 0, 0
	left, right := l1, l2

	for ; left != nil || right != nil || mod != 0; left, right, current = Next(left), Next(right), current.Next {
		leftVal = getVal(left)
		rightVal = getVal(right)
		sum = leftVal + rightVal + mod
		if sum > 9 {
			sum -= 10
			mod = 1
		} else {
			mod = 0
		}
		current.Val = sum
		if Next(left) != nil || Next(right) != nil || mod != 0 {
			current.Next = &ListNode{}
		}
	}

	return out
}

func Next(node *ListNode) *ListNode {
	if node != nil {
		return node.Next
	}
	return nil
}

func getVal(node *ListNode) int {
	if node == nil {
		return 0
	}
	return node.Val
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
