package strivera2zdsa

import "testing"

// https://takeuforward.org/arrays/find-the-missing-number-in-an-array/
func FindMissingNumber(nums []int) int {
	var result int = len(nums) + 1
	for index, value := range nums {
		result ^= value ^ (index + 1)
	}
	return result
}

func TestFindMissingNumber(t *testing.T) {
	tests := []struct {
		array    []int
		expected int
	}{
		{[]int{1, 2, 4, 5}, 3},
		{[]int{1, 2, 3, 4, 5, 6, 7, 9, 10}, 8},
		{[]int{1, 2, 3, 4, 5}, 6},
		{[]int{2, 3, 4, 5, 6, 7}, 1},
		{[]int{2}, 1},
		{[]int{1}, 2},
	}

	for _, test := range tests {
		result := FindMissingNumber(test.array)
		if result != test.expected {
			t.Errorf("findMissingNumber(%v) = %d; expected %d", test.array, result, test.expected)
		}
	}
}
