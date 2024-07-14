package leetcode

func IsSortedOrRotated(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}
	var strikes int
	if nums[0] < nums[len(nums)-1] {
		strikes++
	}
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] > nums[i+1] {
			strikes++
		}

		if strikes >= 2 {
			return false
		}
	}
	return true
}
