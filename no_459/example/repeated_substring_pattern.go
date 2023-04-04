package main

import "strings"

func repeatedSubstringPattern(s string) bool {

	size := len(s)
	sFold := s[1:size] + s[0:size-1]

	return strings.Contains(sFold, s)
}
