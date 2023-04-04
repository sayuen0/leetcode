package main

import "strings"

func lengthOfLastWord(s string) int {
	ss := strings.Split(strings.TrimSpace(s), " ")
	return len(strings.TrimSpace(ss[len(ss)-1]))
}
