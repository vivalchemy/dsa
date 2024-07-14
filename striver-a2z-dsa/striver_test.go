// prime_test.go

package strivera2zdsa

import (
	"reflect"
	"testing"
)

func TestIsPrimeNumber(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},   // Smallest prime number
		{3, true},   // Prime number
		{4, false},  // Not a prime number
		{5, true},   // Prime number
		{10, false}, // Not a prime number
		{13, true},  // Prime number
		{1, false},  // Not a prime number
		{0, false},  // Not a prime number
		{-5, false}, // Negative numbers are not prime
	}

	for _, tc := range testCases {
		result := IsPrimeNumber(tc.input)
		if result != tc.expected {
			t.Errorf("For %d, expected %t, but got %t", tc.input, tc.expected, result)
		}
	}
}

func TestGCD(t *testing.T) {
	testCases := []struct {
		num1     int
		num2     int
		expected int
	}{
		{10, 25, 5},
		{14, 28, 14},
		{21, 7, 7},
		{35, 15, 5},
		{81, 27, 27},
	}

	for _, tc := range testCases {
		got := GCD(tc.num1, tc.num2)
		if got != tc.expected {
			t.Errorf("GCD(%d, %d) = %d; want %d", tc.num1, tc.num2, got, tc.expected)
		}
	}
}

func TestLargestElementInArray(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{"Basic test", []int{1, 2, 3, 4, 5}, 5},
		{"Empty array", []int{}, 0}, // Assuming the function returns 0 for empty array
		{"Single element", []int{99}, 99},
		{"Negative numbers", []int{-1, -2, -3, -4, -5}, -1},
		{"Mixed positive and negative", []int{-5, 1, 0, 10, -2}, 10},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LargestElementInArray(tt.input)
			if result != tt.expected {
				t.Errorf("Expected %d, but got %d for input %v", tt.expected, result, tt.input)
			}
		})
	}
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

// func TestLongestSubarrayWithSumK(t *testing.T) {
// 	testCases := []struct {
// 		N      int
// 		k      int
// 		array  []int
// 		result int
// 	}{
// 		{N: 3, k: 5, array: []int{2, 3, 5}, result: 2},
// 		{N: 3, k: 1, array: []int{-1, 1, 1}, result: 3},
// 		{N: 5, k: 5, array: []int{1, 2, 3, 4, 5}, result: 2},
// 		{N: 5, k: 3, array: []int{1, 1, 1, 1, 1}, result: 3},
// 		{N: 6, k: 0, array: []int{1, -1, 1, -1, 1, -1}, result: 6},
// 		{N: 4, k: -2, array: []int{-1, -1, 2, 1}, result: 2},
// 		{N: 7, k: 7, array: []int{1, 2, 3, 4, 5, 6, 7}, result: 2},
// 		{N: 7, k: 12, array: []int{1, 2, 3, 4, 5, 6, 7}, result: 3},
// 		{N: 4, k: 10, array: []int{1, 2, 3, -2, 6}, result: 5},
// 		{N: 3, k: 15, array: []int{10, 5, -10}, result: 2},
// 	}
//
// 	for _, testCase := range testCases {
// 		output := LongestSubarrayWithGivenSumK(testCase.array, testCase.k)
// 		if output != testCase.result {
// 			t.Errorf("Test failed for input array %v with k=%d. Expected %d, got %d", testCase.array, testCase.k, testCase.result, output)
// 		}
// 	}
// }
