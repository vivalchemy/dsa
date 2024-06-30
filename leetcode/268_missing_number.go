package leetcode

// counting sort
func MissingNumber(nums []int) int {
	var i int
	numLen := len(nums)
	for i < numLen {
		if nums[i] < numLen && nums[i] != nums[nums[i]] {
			nums[i], nums[nums[i]] = nums[nums[i]], nums[i]
		} else {
			i++
		}
	}
	for index, element := range nums {
		if element != index {
			return index
		}
	}
	return numLen
}

func MissingNumberMath(nums []int) int {
	var sum int
	for _, v := range nums {
		sum += v
	}
	return (len(nums)*(len(nums)+1))/2 - sum
}
