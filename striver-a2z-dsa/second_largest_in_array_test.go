package strivera2zdsa

import "testing"

func SecondLargestInArray(arr []int) int {
	if len(arr) <= 1 {
		return -1
	}
	max := arr[0]
	max2 := arr[1]
	if max < max2 && len(arr) == 2 {
		return max
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] > max {
			max2 = max
			max = arr[i]
		} else if arr[i] > max2 {
			max2 = arr[i]
		}
	}
	return max2
}

func TestSecondLargestInArray(t *testing.T) {
	tests := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 1, 1, 1}, 1},
		{[]int{1}, -1}, // or adjust expected result for arrays with less than 2 elements
		{[]int{1, 2}, 1},
		{[]int{-1, -2, -3}, -2},
		{[]int{9, 3, 6, 1, 7, 3}, 7},
		// Add more test cases as needed
	}

	for _, test := range tests {
		result := SecondLargestInArray(test.arr)
		if result != test.expected {
			t.Errorf("For %v, expected %d, but got %d", test.arr, test.expected, result)
		}
	}
}
