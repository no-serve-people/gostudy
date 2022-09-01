package rune

import "testing"

func Test_testRune(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testRune()
		})
	}
}
