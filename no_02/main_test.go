package main

import (
	"github.com/k0kubun/pp/v3"
	"reflect"
	"testing"
)

func Test_addTwoNumbers(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{args: args{
			l1: NewListNode(2, 4, 3),
			l2: NewListNode(5, 6, 4),
		},
			want: NewListNode(7, 0, 8),
		},
		{args: args{
			l1: &ListNode{Val: 0},
			l2: &ListNode{Val: 0},
		},
			want: &ListNode{Val: 0},
		},
		{args: args{
			l1: NewListNode(9, 9, 9, 9, 9, 9, 9),
			l2: NewListNode(9, 9, 9, 9),
		},
			want: NewListNode(8, 9, 9, 9, 0, 0, 0, 1),
		},
		{
			name: "single",
			args: args{
				l1: NewListNode(1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1),
				l2: NewListNode(5, 6, 4),
			},
			want: NewListNode(6, 6, 4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				pp.Println(got)
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewListNode(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
		{args: args{arr: []int{5, 6, 4}}, want: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewListNode(tt.args.arr...); !reflect.DeepEqual(got, tt.want) {
				pp.Println(got)
				t.Errorf("NewListNode() = %v, want %v", got, tt.want)
			}
		})
	}
}
