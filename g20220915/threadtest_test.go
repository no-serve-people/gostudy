package g20220915

import "testing"

func Test_test(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"hehe"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// testChannel()
			testMutex()
		})
	}
}
