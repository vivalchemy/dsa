package strivera2zdsa

import "testing"

func IsSortedOrRotated(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var strikes int
	if nums[0] < nums[len(nums)-1] {
		strikes++
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			strikes++
		}

		if strikes >= 2 {
			return false
		}
	}
	return true
}

func TestIsSortedOrRotated(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"Sorted array", []int{1, 2, 3, 4, 5}, true},
		{"Rotated sorted array", []int{4, 5, 1, 2, 3}, true},
		{"Non-sorted array", []int{2, 3, 1, 4, 5}, false},
		{"Single element", []int{99}, true},
		{"Empty array", []int{}, true}, // Assuming empty array is considered sorted
		{"Descending order array", []int{5, 4, 3, 2, 1}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSortedOrRotated(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %v, but got %v for input %v", tt.expected, result, tt.input)
			}
		})
	}
}
