package strivera2zdsa

import (
	"reflect"
	"testing"
)

func reverse(arr []int) {
	var i int
	arrLen := len(arr) - 1
	for i < arrLen-i {
		arr[i], arr[arrLen-i] = arr[arrLen-i], arr[i]
		i++
	}
}

// NOTE: since no matter the amount of rotation the array will be cyclic
// hence we need to just join the ends the simplest way to do that is to
// reverse the numbers before k and after k in that way the numbers we
// want will be joined but hte array will be in reverse order so we can
// reverse the entire array again to get the answer
func RotateArray(nums []int, k int, isRightShift bool) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}
	k = k % numsLen
	if isRightShift {
		k = numsLen - k
	}
	reverse(nums[0:k])
	reverse(nums[k:numsLen])
	reverse(nums)
}

func TestRotateArray(t *testing.T) {
	tests := []struct {
		nums         []int
		k            int
		isRightShift bool
		expected     []int
	}{
		// Right shifts
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, true, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, true, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3}, 4, true, []int{3, 1, 2}}, // k > len(nums)
		{[]int{1}, 0, true, []int{1}},             // k = 0
		{[]int{1, 2}, 2, true, []int{1, 2}},       // k = len(nums)

		// Left shifts
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, false, []int{4, 5, 6, 7, 1, 2, 3}},
		{[]int{-1, -100, 3, 99}, 2, false, []int{3, 99, -1, -100}},
		{[]int{1, 2, 3}, 4, false, []int{2, 3, 1}}, // k > len(nums)
		{[]int{1}, 0, false, []int{1}},             // k = 0
		{[]int{1, 2}, 2, false, []int{1, 2}},       // k = len(nums)
	}

	for _, test := range tests {
		numsCopy := make([]int, len(test.nums))
		copy(numsCopy, test.nums)
		RotateArray(numsCopy, test.k, test.isRightShift)
		if !reflect.DeepEqual(numsCopy, test.expected) {
			t.Errorf("RotateArray(%v, %d, %v) = %v; expected %v", test.nums, test.k, test.isRightShift, numsCopy, test.expected)
		}
	}
}
