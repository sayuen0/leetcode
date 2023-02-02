package main

import "unicode/utf8"

var n = map[string]int{
	"":   0,
	"I":  1,
	"IV": 4,
	"IX": 9,
	"V":  5,
	"X":  10,
	"XL": 40,
	"XC": 90,
	"L":  50,
	"C":  100,
	"CD": 400,
	"CM": 900,
	"D":  500,
	"M":  1000,
}

func romanToInt(s string) int {
	// 部分文字列を読み込む方法
	if len(s) <= 1 {
		// 1文字以下しかなかったら終了
		return n[s]
	}
	sum := 0

	for i := 0; i < len(s); i++ {
		// 一文字目を読む
		a := sub(s, i)
		// 最後の一文字なら終了
		if len(s) == i+1 {
			sum += n[a]
			break
		}
		// 2文字めを読む
		b := sub(s, i+1)
		// 2文字読んでバリューが見つかったとき、それを採用する
		if c, ok := n[a+b]; ok {
			sum += c
			// さらにカウントを進めておく
			i++
			continue
		}
		// そうでない場合、普通に1文字めを足す
		sum += n[a]
		continue
	}
	return sum
}

func len(s string) int {
	return utf8.RuneCountInString(s)
}

// 数値の文字列を取得
func sub(s string, i int) string {
	return string(s[i])
}
