package main

func plusOne(digits []int) []int {
	out := digits
	mod := 1
	for i := len(digits) - 1; i >= 0; i-- {
		sum := digits[i] + mod
		out[i] = sum % 10
		mod = sum / 10
	}

	if mod == 1 {
		tmp := out
		out = make([]int, len(out)+1)
		copy(out[1:], tmp)
		out[0] = 1
	}
	return out
}
