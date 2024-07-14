package strivera2zdsa

func MoveZeroes(nums []int) {
	numsLen := len(nums)
	if numsLen < 2 {
		return
	}
	var i int
	for j := 0; j < numsLen; j++ {
		if nums[j] != 0 {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
	}

}
