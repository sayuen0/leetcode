package main

import (
	"reflect"
	"testing"
)

func slice(is ...int) []int {
	return is
}

func Test_twoSum(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{args: args{nums: slice(2, 7, 11, 15), target: 9}, want: slice(0, 1)},
		{args: args{nums: slice(3, 2, 4), target: 6}, want: slice(1, 2)},
		{args: args{nums: slice(3, 3), target: 6}, want: slice(0, 1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
