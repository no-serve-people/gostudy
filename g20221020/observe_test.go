package g20221020

import "testing"

func Test_testObserver(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"observer"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testObserver()
		})
	}
}
