package g20221108

import "testing"

func Test_testTemplate(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"template"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testTemplate()
		})
	}
}
