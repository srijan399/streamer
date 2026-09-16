package main

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name     string
		n1, n2   int
		expected int
	}{
		{"first smaller", 3, 7, 3},
		{"second smaller", 10, 2, 2},
		{"equal numbers", 5, 5, 5},
		{"negative numbers", -4, -1, -4},
		{"one negative one positive", -3, 6, -3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMin(tt.n1, tt.n2)
			if result != tt.expected {
				t.Errorf("findMin(%d, %d) = %d; want %d", tt.n1, tt.n2, result, tt.expected)
			}
		})
	}
}
