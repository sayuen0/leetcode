package main

const runeBase = 48

func addStrings(num1 string, num2 string) string {
	if len(num1) < len(num2) {
		num1, num2 = num2, num1
	}
	out := []byte(num1)
	var mod, sum uint8

	for i, j := len(num1)-1, len(num2)-1; i >= 0; i, j = i-1, j-1 {
		a := num1[i] - runeBase
		var b uint8
		if j >= 0 {
			b = num2[j] - runeBase
		}
		sum = a + b + mod
		if sum >= 10 {
			sum -= 10
			mod = 1
		} else {
			mod = 0
		}
		out[i] = sum + runeBase
	}
	if mod == 1 {
		tmp := out
		out = make([]byte, len(out)+1)
		copy(out[1:], tmp)
		out[0] = 1 + runeBase
	}
	return string(out)
}
