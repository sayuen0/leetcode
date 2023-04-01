package main

import (
	"strings"
)

func longestCommonPrefix(s []string) string {
	pref := s[0]
	for i := 1; i < len(s); i++ {
		for !strings.HasPrefix(s[i], pref) {
			pref = pref[:len(pref)-1]
		}
		if pref == "" {
			break
		}
	}
	return pref
}
