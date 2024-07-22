package strivera2zdsa

import (
	// "fmt"
	"reflect"
	"testing"
)

func RearrageElementsBySign(nums []int) []int {
	numsLen := len(nums)
	posCounter := 0
	negCounter := 1
	// fmt.Println("pre nums", nums)
	for posCounter < numsLen && negCounter < numsLen {
		for posCounter < numsLen && nums[posCounter] >= 0 {
			// fmt.Println("nums", nums, "posCounter", posCounter)
			posCounter += 2
		}
		for negCounter < numsLen && nums[negCounter] < 0 {
			// fmt.Println("nums", nums, "negCounter", negCounter)
			negCounter += 2
		}
		if posCounter < numsLen && negCounter < numsLen {
			// fmt.Println("nums", nums, "posCounter", posCounter, "negCounter", negCounter)
			nums[posCounter], nums[negCounter] = nums[negCounter], nums[posCounter]
		}
	}
	return nums
}

func TestRearrageElementsBySign(t *testing.T) {

	testCases := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{1, -1, 2, -2, 3, -3},
			expected: []int{1, -1, 2, -2, 3, -3},
		},
		{
			input:    []int{5, -3, 8, -1, 9, -4},
			expected: []int{5, -3, 8, -1, 9, -4},
		},
		{
			input:    []int{-4, 7, -2, 6},
			expected: []int{7, -4, 6, -2},
		},
		{
			input:    []int{1, -2, 3, -4, 5, -6},
			expected: []int{1, -2, 3, -4, 5, -6},
		},
	}

	for _, tc := range testCases {
		result := RearrageElementsBySign(tc.input)
		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("for input %v, expected %v but got %v", tc.input, tc.expected, result)
		}
	}
}
