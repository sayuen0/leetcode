package main

import (
	"strconv"
)

func addToArrayForm(num []int, k int) []int {
	kStr := strconv.Itoa(k)
	lenNum := len(num)
	lenKStr := len(kStr)
	out := make([]int, Max(lenNum, lenKStr))
	mod := 0
	for i, j := lenNum-1, lenKStr-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		// 配列の方から値を拝借
		a, b := 0, 0
		if i >= 0 {
			a = num[i]
		}
		if j >= 0 {
			b = int(kStr[j] - 48)
		}
		sum := a + b + mod
		if sum >= 10 {
			sum -= 10
			mod = 1
		} else {
			mod = 0
		}
		if i >= j {
			out[i] = sum
		} else {
			out[j] = sum
		}
	}
	if mod == 1 {
		tmp := out
		out = make([]int, len(out)+1)
		copy(out[1:], tmp)
		out[0] = 1
	}
	return out
}

func Max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
