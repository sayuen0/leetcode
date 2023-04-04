package main

import "strings"

func lengthOfLastWord(s string) int {
	str := strings.Fields(s) //tokenize
	return len(str[len(str)-1])
}
