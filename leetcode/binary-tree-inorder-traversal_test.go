package leetcode

import (
	"reflect"
	"testing"
)

func Test_inorderTraversal(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name    string
		args    args
		wantAns []int
	}{
		{"1", args{&TreeNode{Val: 1, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}}}}, []int{1, 3, 2}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := inorderTraversal(tt.args.root); !reflect.DeepEqual(gotAns, tt.wantAns) {
				t.Errorf("inorderTraversal() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
