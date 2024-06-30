package leetcode

// when you can't modify the original array
// counting sort
func FindDuplicate(nums []int) int {
	numsLen := len(nums)
	dummySlice := make([]int, numsLen-1)
	for _, value := range nums {
		if dummySlice[value-1] == value {
			return value
		}
		dummySlice[value-1] = value
	}
	return 0
}

// when  you can modify the original array
func FindDuplicateByNegation(nums []int) int {
	for _, val := range nums {
		index := abs(val) - 1
		if nums[index] < 0 {
			return index + 1
		}
		nums[index] *= -1
	}
	return -1
}
