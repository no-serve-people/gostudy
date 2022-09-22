package g20220922

import "testing"

func Test_deferCall3(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deferCall3()
		})
	}
}

func Test_deferCall(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			deferCall()
		})
	}
}
