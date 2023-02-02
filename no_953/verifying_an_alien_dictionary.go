package main

import "strings"

func isAlienSorted(words []string, order string) bool {
	// 配列長が1の場合終了
	if len(words) == 1 {
		return true
	}

	l := len(words)
	for i, j := 0, 1; i < l-1; i, j = i+1, j+1 {
		left, right := words[i], words[j]
		lenL := len(left)  // word
		lenR := len(right) // world

		pairOrdered := false
		for k := 0; k < lenL && k < lenR; k++ {

			leftW := left[k]   // w
			rightW := right[k] // w

			// w > w
			p, q := strings.Index(order, string(leftW)), strings.Index(order, string(rightW))
			if p > q {
				return false
			} else if p < q {
				pairOrdered = true
				break
			}

			// どちらかの最後のインデックスに達したか
			// ex: kuvp, p ,ngxlkthsjuoqcpavbfdermiywz
			if k == lenL-1 || k == lenR-1 {
				break
			}
		}

		// both have same prefix "app" but different length
		// len(apple) > len(app)
		if !pairOrdered && lenL > lenR {
			return false
		}
	}
	return true
}

func Max(a, b int) (max int, left bool) {
	if a >= b {
		return a, true
	}
	return b, false
}
