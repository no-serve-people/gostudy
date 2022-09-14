package g20220914

import "testing"

func Test_testInterface(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"test"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testInterface()
		})
	}
}
