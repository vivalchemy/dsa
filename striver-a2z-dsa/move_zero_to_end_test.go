package strivera2zdsa

import (
	"reflect"
	"testing"
)

func MoveZeroes(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}
	var i int
	for j := 0; j < numsLen; j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

}

func TestMoveZeroes(t *testing.T) {
	tests := []struct {
		nums     []int
		expected []int
	}{
		{[]int{0, 1, 0, 3, 12}, []int{1, 3, 12, 0, 0}},
		{[]int{0, 0, 1}, []int{1, 0, 0}},
		{[]int{0}, []int{0}},
		{[]int{1, 0}, []int{1, 0}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{0, 0, 0, 0, 1}, []int{1, 0, 0, 0, 0}},
		{[]int{}, []int{}}, // empty slice
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		MoveZeroes(numsCopy)
		if !reflect.DeepEqual(numsCopy, test.expected) {
			t.Errorf("MoveZeroes(%v) = %v; expected %v", test.nums, numsCopy, test.expected)
		}
	}
}
