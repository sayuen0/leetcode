package main

func addBinary(a, b string) string {
	left, right := []rune(a), []rune(b)
	out := ""
	carry := '0'
	// 逆順から読み取る
	lenLeft, lenRight := len(left), len(right)
	for i, j := lenLeft-1, lenRight-1; i >= 0 || j >= 0; i, j = i-1, j-1 {
		// 左をとる。配列長を超えてたら0
		l, r := '0', '0'
		if i >= 0 {
			l = left[i]
		}
		// 右をとる。配列長を超えてたら0
		if j >= 0 {
			r = right[j]
		}
		var sum = '0'
		sum, carry = FullAdder(l, r, carry)
		out = string(sum) + out
	}
	if carry == '1' {
		out = string(carry) + out
	}
	return out
}

// FullAdder 全加算器
func FullAdder(l, r, modIn rune) (sum, modOut rune) {
	carry := 0
	if l == '1' {
		carry++
	}
	if r == '1' {
		carry++
	}
	if modIn == '1' {
		carry++
	}
	switch carry {
	case 0:
		return '0', '0'
	case 1:
		return '1', '0'
	case 2:
		return '0', '1'
	default:
		return '1', '1'
	}
}
