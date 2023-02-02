package main

import "testing"

func Test_isAlienSorted(t *testing.T) {
	type args struct {
		words []string
		order string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"1", args{words: []string{"hello", "leetcode"}, order: "hlabcdefgijkmnopqrstuvwxyz"}, true},
		{"2", args{words: []string{"word", "world", "row"}, order: "worldabcefghijkmnpqstuvxyz"}, false},
		{"3", args{words: []string{"apple", "app"}, order: "abcdefghijklmnopqrstuvwxyz"}, false},
		{"58", args{words: []string{"kuvp", "q"}, order: "ngxlkthsjuoqcpavbfdermiywz"}, true},
		{"78", args{words: []string{"fxasxpc", "dfbdrifhp", "nwzgs", "cmwqriv", "ebulyfyve", "miracx", "sxckdwzv", "dtijzluhts", "wwbmnge", "qmjwymmyox"}, order: "zkgwaverfimqxbnctdplsjyohu"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAlienSorted(tt.args.words, tt.args.order); got != tt.want {
				t.Errorf("isAlienSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}
