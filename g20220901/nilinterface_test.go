package g20220901

import "testing"

func Test_testNilInterface2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testNilInterface2()
		})
	}
}

func Test_testNilInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testNilInterface()
		})
	}
}
