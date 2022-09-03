package leetcode

import "testing"

func Test_findLongestChain(t *testing.T) {
	type args struct {
		pairs [][]int
	}
	tests := []struct {
		name    string
		args    args
		wantAns int
	}{
		{"1", args{[][]int{{1, 2}, {2, 3}, {3, 4}}}, 2},
		{"2", args{[][]int{{1, 2}, {7, 8}, {4, 5}}}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := findLongestChain(tt.args.pairs); gotAns != tt.wantAns {
				t.Errorf("findLongestChain() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
