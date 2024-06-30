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
