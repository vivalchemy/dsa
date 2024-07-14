package strivera2zdsa

import "testing"

func LinearSearch(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}
	return -1
}

func TestLinearSearch(t *testing.T) {
	tests := []struct {
		nums     []int
		target   int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 3, 2},
		{[]int{1, 2, 3, 4, 5}, 1, 0},
		{[]int{1, 2, 3, 4, 5}, 5, 4},
		{[]int{1, 2, 3, 4, 5}, 0, -1},
		{[]int{5, 4, 3, 2, 1}, 4, 1},
		{[]int{5, 4, 3, 2, 1}, 6, -1},
		{[]int{}, 1, -1}, // empty slice
	}

	for _, test := range tests {
		result := LinearSearch(test.nums, test.target)
		if result != test.expected {
			t.Errorf("LinearSearch(%v, %d) = %d; expected %d", test.nums, test.target, result, test.expected)
		}
	}
}
