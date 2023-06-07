package util

import "testing"

func TestModLikePython(t *testing.T) {
	tests := []struct {
		name string
		n    int
		m    int
		want int
	}{
		{"Test 1", 10, 3, 1},
		{"Test 2", -10, 3, 2},
		{"Test 3", 10, -3, 1},
		{"Test 4", -10, -3, 2},
		{"Test 5", 0, 1, 0},
		{"Test 6", 10, 1, 0},
		{"Test 7", 0, -1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModLikePython(tt.n, tt.m); got != tt.want {
				t.Errorf("ModLikePython() = %v, want %v", got, tt.want)
			}
		})
	}
}
