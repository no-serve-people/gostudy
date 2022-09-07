package leetcode

import "testing"

func Test_reorderSpaces(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		wantAns string
	}{
		{"1", args{"  this   is  a sentence "}, "this   is   a   sentence"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAns := reorderSpaces(tt.args.s); gotAns != tt.wantAns {
				t.Errorf("reorderSpaces() = %v, want %v", gotAns, tt.wantAns)
			}
		})
	}
}
