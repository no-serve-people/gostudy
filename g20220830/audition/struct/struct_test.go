package _struct

import "testing"

func Test_testStruct(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testStruct()
		})
	}
}
