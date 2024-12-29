package leetcode

import (
	"reflect"
	"testing"
)

func sortArray(nums []int){
  for i := 0; i < len(nums)-1; i++ {
    for j := i; j < len(nums); j++ {
      if nums[i] > nums[j] {
        nums[i], nums[j] = nums[j], nums[i]
      }
    }
  }
}


// get the array
// make another array with the same length as the input array
// sort them
// store the location in the sorted array in the newly created array
// return the result

// if there are repeated numbers then we can't use the above solution as
// 1 2 2 2 3
//       ^
//       |_ for this the output should be 1 but it will be 3
// 1 2 2 2 3
//         ^
//         |_ for 3 the output should be 4 which will be correct
// hence we can't use this solution
func HowManyNumsSmallerThanCurrentNum(nums []int)[]int {
  result := make([]int, len(nums))

  arrayForSorting := make([]int, len(nums))
  copy(arrayForSorting, nums)
  sortArray(arrayForSorting)

  tmpMap := make(map[int]int)
  for index, val := range arrayForSorting{
    if _, ok := tmpMap[val]; !ok {
      tmpMap[val] = index
    }
  }

  for index, val := range nums{
      result[index] = tmpMap[val]
  }
  return result
}

func TestHowManyNumsSmallerThanCurrentNum(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{
			input:    []int{8, 1, 2, 2, 3},
			expected: []int{4, 0, 1, 1, 3},
		},
		{
			input:    []int{6, 5, 4, 8},
			expected: []int{2, 1, 0, 3},
		},
		{
			input:    []int{7, 7, 7, 7},
			expected: []int{0, 0, 0, 0},
		},
	}

	for _, test := range tests {
		result := HowManyNumsSmallerThanCurrentNum(test.input)
		if !reflect.DeepEqual(result, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, result)
		}
	}
}
