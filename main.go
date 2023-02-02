package main

import (
	"fmt"
	"unicode/utf8"
)

func Reverse(s string) string {
	a := ""
	for i := len(s) - 1; i >= 0; i-- {
		a += string(s[i])
	}
	return a
}

func ReverseBySwap(s string) string {
	runes := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func ReverseMultiByteBySwap(s string) string {
	runes := []rune(s)
	for i, j := 0, utf8.RuneCountInString(s)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(ReverseBySwap("Hello World"))
	fmt.Println(ReverseMultiByteBySwap("こんにちは世界"))
}
