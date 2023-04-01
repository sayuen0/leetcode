package main

import "testing"

func Test_search(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
		{"", args{[]int{-1, 0, 3, 5, 9, 12}, 9}, 4},
		{"", args{[]int{-1, 0, 3, 5, 9, 12}, 2}, -1},
		{"", args{[]int{2, 5}, 2}, 0},
		{"24", args{[]int{-1, 0, 5}, -1}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
