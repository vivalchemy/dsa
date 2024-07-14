package leetcode

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return 1
	}
	swapLocationPointer := 0
	for currentIndex := 1; currentIndex < len(nums); currentIndex++ {
		if nums[currentIndex] != nums[swapLocationPointer] {
			swapLocationPointer++
			nums[swapLocationPointer] = nums[currentIndex]
		}
	}

	return swapLocationPointer + 1
}
