package g20220901

import "testing"

func Test_testPanic2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testPanic2()
		})
	}
}
