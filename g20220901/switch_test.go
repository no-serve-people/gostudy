package g20220901

import "testing"

func Test_testSwitch(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSwitch()
		})
	}
}
