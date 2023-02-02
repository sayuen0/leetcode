package main

import "unicode/utf8"

var m = map[rune]int{
	'0': 0,
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func _romanToInt(s string) int {
	runeBefore := rune(0)
	sum := 0
	for i, r := range s {
		switch r {
		// I, X, Cが来た時、次がVかどうかチェックする
		case 'I':
			// Iは組み合わせて何かになるので、足さない
			// ただし、ラストなら足す
			if last(s, i) {
				sum += m[r]
				break
			}
		case 'V':
			// 先にIが入っていたら、4
			if runeBefore == 'I' {
				sum += 4
				break
			}
			// それ以外はそのまま足す
			sum += m[r]
		case 'X':
			// 先にIが入っていたら、9
			if runeBefore == 'I' {
				sum += 9
				break
			}

			// ここに来た場合、直前の値は足されているので考慮不要
			// Xは次のと合わさって何かになるので足さない
			// ただし、最後の一文字だった場合はそのまま足す
			if last(s, i) {
				sum += m[r]
				break
			}

		case 'L':
			// 先にXが入っていたら、40
			if runeBefore == 'X' {
				sum += 40
				break
			}

			// それ以外はそのままたす
			sum += m[r]
		case 'C':
			// 先にXが入っていたら、90
			if runeBefore == 'X' {
				sum += 90
				break
			}

			// Cは合わさって何かになるので足さない
			// ただし、最後の一文字だった場合はそのまま足す
			if last(s, i) {
				sum += m[r]
				break
			}
		case 'D':
			// 先にCが入っていたら、400
			if runeBefore == 'C' {
				sum += 400
				break
			}

			// それ以外はそのままたす
			sum += m[r]

		case 'M':
			// 先にCが入っていたら、600
			if runeBefore == 'C' {
				sum += 900
				break
			}

			// それ以外はそのままたす
			sum += m[r]

		}
		// 次の週でのチェックのため今の週のルーンを入れる
		runeBefore = r
	}
	return sum
}

func last(s string, i int) bool {
	return utf8.RuneCountInString(s)-1 == i
}
