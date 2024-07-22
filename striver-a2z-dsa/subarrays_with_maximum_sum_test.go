package strivera2zdsa

import (
	// "fmt"
	"reflect"
	"testing"
)

func SubarraysWithMaximumSum(nums []int) (int, [][2]int) {
	numsLen := len(nums)
	// fmt.Println(nums)
	if numsLen <= 0 {
		return 0, [][2]int{}
	}
	max := nums[0]
	solnArr := make([][2]int, 0)
	sum := 0
	var start int
	for index := 0; index < numsLen; index++ {
		value := nums[index]
		sum += value
		if max == sum {
			// fmt.Println("Appending ", [2]int{start, index})
			solnArr = append(solnArr, [2]int{start, index})
		} else if max < sum {
			solnArr = [][2]int{}
			max = sum
			// fmt.Println("Appending ", [2]int{start, index})
			solnArr = append(solnArr, [2]int{start, index})
		}
		if sum <= 0 {
			sum = 0
			// fmt.Println("Changing start from ", start, " to ", index+1)
			start = index + 1
		}
	}
	return max, solnArr
}

func TestSubarraysWithMaximumSum(t *testing.T) {
	tests := []struct {
		arr             []int
		expectedSum     int
		expectedIndices [][2]int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6, [][2]int{{3, 6}}},
		{[]int{1}, 1, [][2]int{{0, 0}}},
		{[]int{5, 4, -1, 7, 8}, 23, [][2]int{{0, 4}}},
		{[]int{-1, -2, -3, -4}, -1, [][2]int{{0, 0}}},
		{[]int{1, 2, 3, 4, 5}, 15, [][2]int{{0, 4}}},
		{[]int{}, 0, [][2]int{}},
		{[]int{0, -3, 1, 2, -1, 2, -1, 3, -2, 3}, 7, [][2]int{{2, 9}}},
		{[]int{3, 3, -2, 3, 3}, 10, [][2]int{{0, 4}}},
		{[]int{3, -2, 3, 3}, 7, [][2]int{{0, 3}}},
		{[]int{1, -2, 1, -2, 1, -2, 1}, 1, [][2]int{{0, 0}, {2, 2}, {4, 4}, {6, 6}}},
		{[]int{2, -1, 2, -1, 2, -1, 2}, 5, [][2]int{{0, 6}}},
		{[]int{-1, 2, 3, -4, 2, 3, -4, 2, 3}, 7, [][2]int{{1, 8}}},
		{[]int{1, 2, -3, 1, 2, -3, 1, 2}, 3, [][2]int{{0, 1}, {3, 4}, {6, 7}}},
		{[]int{4, -1, 4, -1, 4, -1, 4}, 13, [][2]int{{0, 6}}},
	}

	for _, test := range tests {
		sum, indices := SubarraysWithMaximumSum(test.arr)
		if sum != test.expectedSum || !reflect.DeepEqual(indices, test.expectedIndices) {
			t.Errorf("For array %v, expected sum %d and subarrays %v, but got sum %d and subarrays %v",
				test.arr, test.expectedSum, test.expectedIndices, sum, indices)
		}
	}
}
