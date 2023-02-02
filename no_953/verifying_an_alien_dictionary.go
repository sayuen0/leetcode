package main

import "strings"

func isAlienSorted(words []string, order string) bool {
	// 配列長が1の場合終了
	l := len(words)
	if l == 1 {
		return true
	}

	for i, j := 0, 1; i < l-1; i, j = i+1, j+1 {
		left, right := words[i], words[j]
		if !dictionaryOrdered(left, right, order) {
			return false
		}
	}

	return true
}

func dictionaryOrdered(left string, right string, order string) (ordered bool) {
	lenL := len(left)  // word
	lenR := len(right) // world

	for k := 0; k < lenL && k < lenR; k++ {
		leftW := left[k]   // d
		rightW := right[k] // l

		p, q := strings.Index(order, string(leftW)), strings.Index(order, string(rightW))
		// d != l
		if p != q {
			return p < q
		}
	}

	return lenL <= lenR
}
