package strivera2zdsa

import "testing"

func LongestSubarrayWithGivenSumK(nums []int, target int) int {
	var maxLength int
	var sum int
	sumMap := make(map[int]int)
	sumMap[0] = -1
	for index, num := range nums {
		sum += num
		if startIndex, exists := sumMap[sum-target]; (exists) && (maxLength < index-startIndex) {
			maxLength = index - startIndex
		}
		if _, exists := sumMap[sum]; !exists {
			sumMap[sum] = index
		}
	}
	return maxLength
}

func TestLongestSubarrayWithSumK(t *testing.T) {
	testCases := []struct {
		N      int
		k      int
		array  []int
		result int
	}{
		{N: 3, k: 5, array: []int{2, 3, 5}, result: 2},
		{N: 3, k: 1, array: []int{-1, 1, 1}, result: 3},
		{N: 5, k: 5, array: []int{1, 2, 3, 4, 5}, result: 2},
		{N: 5, k: 3, array: []int{1, 1, 1, 1, 1}, result: 3},
		{N: 6, k: 0, array: []int{1, -1, 1, -1, 1, -1}, result: 6},
		{N: 4, k: -2, array: []int{-1, -1, 2, 1}, result: 2},
		{N: 7, k: 7, array: []int{1, 2, 3, 4, 5, 6, 7}, result: 2},
		{N: 7, k: 12, array: []int{1, 2, 3, 4, 5, 6, 7}, result: 3},
		{N: 4, k: 10, array: []int{1, 2, 3, -2, 6}, result: 5},
		{N: 3, k: 15, array: []int{10, 5, -10}, result: 2},
	}

	for _, testCase := range testCases {
		output := LongestSubarrayWithGivenSumK(testCase.array, testCase.k)
		if output != testCase.result {
			t.Errorf("Test failed for input array %v with k=%d. Expected %d, got %d", testCase.array, testCase.k, testCase.result, output)
		}
	}
}
