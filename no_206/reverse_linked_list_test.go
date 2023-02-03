package main

import (
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{"1", args{head: NewListNode(1, 2, 3, 4, 5)}, NewListNode(5, 4, 3, 2, 1)},
		{"2", args{head: NewListNode(1, 2)}, NewListNode(2, 1)},
		{"3", args{head: NewListNode()}, NewListNode()},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
