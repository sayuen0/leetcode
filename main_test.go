package main

import (
	"fmt"
	"testing"
)

func FuzzReverse(f *testing.F) {
	f.Fuzz(func(t *testing.T, s string) {
		fmt.Println(Reverse(s))
	})
}

const word = "Lorem Ipsum is simply dummy text of the printing and typesetting industry. Lorem Ipsum has been the industry's standard dummy text ever since the 1500s, when an unknown printer took a galley of type and scrambled it to make a type specimen book. It has survived not only five centuries, but also the leap into electronic typesetting, remaining essentially unchanged. It was popularised in the 1960s with the release of Letraset sheets containing Lorem Ipsum passages, and more recently with desktop publishing software like Aldus PageMaker including versions of Lorem Ipsum."

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(word)
	}
}

func BenchmarkReverseBySwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBySwap(word)
	}
}

func BenchmarkMultiByteReverseBySwap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ReverseBySwap(word)
	}
}
