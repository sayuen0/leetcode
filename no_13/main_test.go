package main

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{args: args{s: "III"}, want: 3},
		{args: args{s: "LVIII"}, want: 50 + 5 + 3},
		{args: args{s: "MCMXCIV"}, want: 1000 + 900 + 90 + 4},
		{args: args{s: "IX"}, want: 9},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
