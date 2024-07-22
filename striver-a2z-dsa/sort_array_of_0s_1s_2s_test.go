package strivera2zdsa

import (
	"fmt"
	"reflect"
	"testing"
)

// with the help of counters
// func SortArrayOf0s1s2s(nums []int) []int {
// 	zerosPointer := 0
// 	twosPointer := len(nums)
// 	fmt.Println("pre num", nums)
// 	for _, value := range nums {
// 		if value == 0 {
// 			zerosPointer++
// 		} else if value == 2 {
// 			twosPointer--
// 		}
// 	}
// 	fmt.Println("post calc", nums, zerosPointer, twosPointer)
// 	for i := 0; i < zerosPointer; i++ {
// 		nums[i] = 0
// 	}
// 	fmt.Println("post 0", nums)
// 	for i := zerosPointer; i < twosPointer; i++ {
// 		nums[i] = 1
// 	}
// 	fmt.Println("post 1", nums)
// 	for i := twosPointer; i < len(nums); i++ {
// 		nums[i] = 2
// 	}
// 	fmt.Println("post 2", nums)
// 	return nums
// }

func SortArrayOf0s1s2s(nums []int) []int {
	zerosPointer := 0
	twosPointer := len(nums) - 1
	fmt.Println("pre num", nums)
	for index := 0; index <= twosPointer; index++ {
		fmt.Println("nums", nums, "index", index, "zerosPointer", zerosPointer, "twosPointer", twosPointer)
		value := nums[index]
		if value == 0 {
			// index is not decremented since the possible received values are only 1
			nums[zerosPointer], nums[index] = 0, nums[zerosPointer]
			zerosPointer++
		} else if value == 2 {
			nums[twosPointer], nums[index] = 2, nums[twosPointer]
			twosPointer--
			// index is decremented since the possible values are 0, 1, 2 so we need to apply the conditions on the received values again
			index--
		}
	}
	fmt.Println("post num", nums, zerosPointer, twosPointer)
	return nums
}

func TestSortArrayOf0s1s2s(t *testing.T) {
	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{2, 0, 2, 1, 1, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{2, 0, 1}, []int{0, 1, 2}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
		{[]int{0, 2, 1}, []int{0, 1, 2}},
		{[]int{2, 2, 1, 1, 0, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{1, 0, 2, 1, 2, 0}, []int{0, 0, 1, 1, 2, 2}},
		{[]int{1, 1, 0, 0, 2, 2}, []int{0, 0, 1, 1, 2, 2}},
	}

	for _, test := range tests {
		SortArrayOf0s1s2s(test.input)
		if !reflect.DeepEqual(test.input, test.expected) {
			t.Errorf("For input %v, expected %v but got %v", test.input, test.expected, test.input)
		}
	}
}
