package sol

import (
	"reflect"
	"testing"
)

func Test_mergeKLists(t *testing.T) {
	type args struct {
		lists []*ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "Example 1",
			args: args{
				lists: []*ListNode{
					BuildListFromArray([]int{1, 4, 5}),
					BuildListFromArray([]int{1, 3, 4}),
					BuildListFromArray([]int{2, 6}),
				},
			},
			want: BuildListFromArray([]int{1, 1, 2, 3, 4, 4, 5, 6}),
		},
		{
			name: "Example 2",
			args: args{
				lists: []*ListNode{},
			},
			want: BuildListFromArray([]int{}),
		},
		{
			name: "Example 3",
			args: args{
				lists: []*ListNode{
					BuildListFromArray([]int{}),
				},
			},
			want: BuildListFromArray([]int{}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeKLists(tt.args.lists); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeKLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BuildListFromArray(arr []int) *ListNode {
	var head *ListNode
	var current *ListNode
	for _, val := range arr {
		node := &ListNode{Val: val}
		if head == nil {
			head = node
			current = node
		} else {
			current.Next = node
			current = current.Next
		}
	}
	return head
}
