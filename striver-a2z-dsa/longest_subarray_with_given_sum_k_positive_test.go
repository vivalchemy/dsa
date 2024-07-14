package strivera2zdsa

import "testing"

func LongestSubarrayWithGivenSumKPositive(nums []int, target int) int {
	var backPointer, maxLen, sum int
	// fmt.Println("---------Start---------")
	// fmt.Println("Given Array : ", nums, "\tTarget : ", target)
	for frontPointer := range nums {
		sum += nums[frontPointer]
		// fmt.Println("fp", frontPointer, "\tbp", backPointer, "\tSum", sum, "\tlen", frontPointer-backPointer+1, "\tMax", maxLen)
		for (sum > target) && (backPointer < frontPointer) {
			sum -= nums[backPointer]
			backPointer++
			// fmt.Println("Changing bp: fp", frontPointer, "\tbp", backPointer, "\tSum", sum, "\tlen", frontPointer-backPointer+1, "\tMax", maxLen)
		}
		if (sum == target) && (frontPointer-backPointer+1 > maxLen) {
			maxLen = frontPointer - backPointer + 1
		}
	}
	// fmt.Println("Max Length : ", maxLen)
	// fmt.Println("---------End--------\n-")
	return maxLen
}

func TestLongestSubarrayWithSumKPositive(t *testing.T) {
	testCases := []struct {
		arr      []int
		k        int
		expected int
	}{
		{[]int{2, 3, 5}, 5, 2},
		{[]int{2, 3, 5, 1, 9}, 10, 3},
		{[]int{1, 2, 3, 7, 5}, 12, 3},
		{[]int{1, 2, 3, 4, 5}, 15, 5},
		{[]int{10, 5, 2, 7, 1, 9}, 15, 4},
		{[]int{0, 0, 0, 0, 0}, 0, 5},
		{[]int{0, 0, 0, 0, 5}, 5, 5},
		{[]int{1, 0, 3, 0, 5}, 4, 4},
	}

	for _, tc := range testCases {
		result := LongestSubarrayWithGivenSumKPositive(tc.arr, tc.k)
		if result != tc.expected {
			t.Errorf("For input arr = %v, k = %d, expected %d but got %d", tc.arr, tc.k, tc.expected, result)
		}
	}
}
