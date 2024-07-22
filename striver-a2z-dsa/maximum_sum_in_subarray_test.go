package strivera2zdsa

import (
	// "math"
	"reflect"
	"testing"
)

// if we don't want to use the math library
// func MaximumSumInSubarray(nums []int) (int, []int) {
// 	if len(nums) <= 0 {
// 		return 0, []int{}
// 	}
// 	max := math.MinInt32
// 	var start, maxStart, maxEnd, sum int
// 	for index, value := range nums {
// 		sum += value
// 		if max < sum {
// 			max = sum
// 			maxStart = start
// 			maxEnd = index
// 		}
// 		if sum <= 0 {
// 			sum = 0
// 			start = index + 1
// 		}
// 	}
// 	return max, nums[maxStart : maxEnd+1]
// }

func MaximumSumInSubarray(nums []int) (int, []int) {
	numsLen := len(nums)
	if numsLen <= 0 {
		return 0, []int{}
	}
	max := nums[0]
	sum := 0
	maxEnd := 0
	var start, maxStart int
	for index := 0; index < numsLen; index++ {
		value := nums[index]
		sum += value
		if max < sum {
			max = sum
			maxStart = start
			maxEnd = index
		}
		if sum <= 0 {
			sum = 0
			start = index + 1
		}
	}
	return max, nums[maxStart : maxEnd+1]
}

func TestMaximumSumInSubarray(t *testing.T) {
	tests := []struct {
		arr              []int
		expectedSum      int
		expectedSubarray []int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, []int{4, -1, 2, 1}},
		{[]int{1}, 1, []int{1}},
		{[]int{5, 4, -1, 7, 8}, 23, []int{5, 4, -1, 7, 8}},
		{[]int{-1, -2, -3, -4}, -1, []int{-1}},
		{[]int{1, 2, 3, 4, 5}, 15, []int{1, 2, 3, 4, 5}},
		{[]int{}, 0, []int{}},
		{[]int{0, -3, 1, 2, -1, 2, -1, 3, -2, 3}, 7, []int{1, 2, -1, 2, -1, 3, -2, 3}},
		{[]int{-2, -1, -4}, -1, []int{-1}},
	}

	for _, test := range tests {
		sum, subarray := MaximumSumInSubarray(test.arr)
		if sum != test.expectedSum || !reflect.DeepEqual(subarray, test.expectedSubarray) {
			t.Errorf("For array %v, expected sum %d and subarray %v, but got sum %d and subarray %v",
				test.arr, test.expectedSum, test.expectedSubarray, sum, subarray)
		}
	}
}
