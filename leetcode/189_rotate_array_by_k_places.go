package leetcode

func reverse(arr []int) {
	var i int
	arrLen := len(arr) - 1
	for i < arrLen-i {
		arr[i], arr[arrLen-i] = arr[arrLen-i], arr[i]
		i++
	}
}

// NOTE: since no matter the amount of rotation the array will be cyclic
// hence we need to just join the ends the simplest way to do that is to
// reverse the numbers before k and after k in that way the numbers we
// want will be joined but hte array will be in reverse order so we can
// reverse the entire array again to get the answer
func RotateArray(nums []int, k int, isRightShift bool) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}
	k = k % numsLen
	if isRightShift {
		k = numsLen - k
	}
	reverse(nums[0:k])
	reverse(nums[k:numsLen])
	reverse(nums)
}
