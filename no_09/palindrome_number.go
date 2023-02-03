package main

import "strconv"

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	l := len(s)
	middle := l / 2
	for i, j := 0, l-1; i <= middle-1; i, j = i+1, j-1 {
		if !sameLetter(s[i], s[j]) {
			return false
		}
	}
	return true
}

func sameLetter(left, right uint8) bool {
	return left == right
}
