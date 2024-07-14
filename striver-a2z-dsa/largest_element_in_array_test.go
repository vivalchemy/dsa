package strivera2zdsa

import "testing"

func LargestElementInArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func TestLargestElementInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Basic test", []int{1, 2, 3, 4, 5}, 5},
		{"Empty array", []int{}, 0}, // Assuming the function returns 0 for empty array
		{"Single element", []int{99}, 99},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -1},
		{"Mixed positive and negative", []int{-5, 1, 0, 10, -2}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestElementInArray(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d for input %v", tt.expected, result, tt.input)
			}
		})
	}
}
