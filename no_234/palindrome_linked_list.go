package main

func isPalindrome(head *ListNode) bool {
	m := make(map[int]int, 0)
	listIndex := 0
	current := head
	for {
		m[listIndex] = current.Val
		if !hasNext(current) {
			break
		}
		current = current.Next
		listIndex++
	}

	lenM := len(m)
	for i, j := 0, lenM-1; i < lenM/2; i, j = i+1, j-1 {
		if !sameNum(m[i], m[j]) {
			return false
		}
	}
	return true
}

func hasNext(node *ListNode) bool {
	return node.Next != nil
}

func sameNum(a, b int) bool {
	return a == b
}
