package g20220901

import "testing"

func Test_testArray(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArray()
		})
	}
}

func Test_testArray2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"2"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testArray2()
		})
	}
}

func Test_testSlice1(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testSlice1()
		})
	}
}

func Test_testNilSlice2(t *testing.T) {
	tests := []struct {
		name string
	}{
		{"3"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testNilSlice2()
		})
	}
}
