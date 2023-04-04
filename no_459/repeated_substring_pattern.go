package main

func repeatedSubstringPattern(s string) bool {
	runes := []rune(s)
	l := len(runes)
	for i := 1; i < l; i++ {
		subs := runes[0:i]
		eq := true
		for j := i; j < l; j += len(subs) {
			target := runes[j : j+len(subs)]
			if !equal(target, subs) {
				eq = false
				break
			}
		}
		if eq {
			return true
		}
	}
	return false
}

func equal(a, b []rune) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
