package strivera2zdsa

import (
	"reflect"
	"testing"
)

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}
	swapLocationPointer := 0
	for currentIndex := 1; currentIndex < len(nums); currentIndex++ {
		if nums[currentIndex] != nums[swapLocationPointer] {
			swapLocationPointer++
			nums[swapLocationPointer] = nums[currentIndex]
		}
	}

	return swapLocationPointer + 1
}

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name        string
		input       []int
		expectedK   int
		expectedArr []int // This will represent the expected unique elements in the array up to `expectedK` elements
	}{
		{"Example 1", []int{1, 1, 2}, 2, []int{1, 2}},
		{"Example 2", []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}, 5, []int{0, 1, 2, 3, 4}},
		{"Empty array", []int{}, 0, []int{}},
		{"Single element", []int{99}, 1, []int{99}},
		{"All duplicates", []int{1, 1, 1, 1, 1}, 1, []int{1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Make a copy of the input slice to avoid modifying the original
			inputCopy := make([]int, len(tt.input))
			copy(inputCopy, tt.input)

			k := RemoveDuplicates(inputCopy)

			// Check if k matches expected number of unique elements
			if k != tt.expectedK {
				t.Errorf("Expected %d unique elements, but got %d for input %v", tt.expectedK, k, tt.input)
			}

			// Check if the first k elements of inputCopy match the expected unique elements
			if !reflect.DeepEqual(inputCopy[:k], tt.expectedArr) {
				t.Errorf("After removing duplicates, expected %v, but got %v for input %v", tt.expectedArr, inputCopy[:k], tt.input)
			}
		})
	}
}
