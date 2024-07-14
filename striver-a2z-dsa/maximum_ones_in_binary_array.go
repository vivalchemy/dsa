package strivera2zdsa

// // counter approach
// func MaximumOnes(nums []int) int {
// 	var max, current int
// 	for _, value := range nums {
// 		if value == 1 {
// 			current++
// 		} else if current > max {
// 			max = current
// 			current = 0
// 		} else {
// 			current = 0
// 		}
// 	}
// 	if current > max {
// 		return current
// 	}
// 	return max
// }

// sliding window approach
func MaximumOnes(nums []int) int {
	var winStart, winEnd, curWinSize, max int
	for winEnd < len(nums) {
		if nums[winEnd] == 0 {
			curWinSize = winEnd - winStart
			if curWinSize > max {
				max = curWinSize
			}
			winStart = winEnd + 1
		}
		winEnd++
	}
	if winEnd-winStart > max {
		max = winEnd - winStart
	}
	return max
}
