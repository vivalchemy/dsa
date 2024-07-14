package strivera2zdsa

func LongestSubarrayWithGivenSumKPositive(nums []int, target int) int {
	var backPointer, maxLen, sum int
	// fmt.Println("---------Start---------")
	// fmt.Println("Given Array : ", nums, "\tTarget : ", target)
	for frontPointer := range nums {
		sum += nums[frontPointer]
		// fmt.Println("fp", frontPointer, "\tbp", backPointer, "\tSum", sum, "\tlen", frontPointer-backPointer+1, "\tMax", maxLen)
		for (sum > target) && (backPointer < frontPointer) {
			sum -= nums[backPointer]
			backPointer++
			// fmt.Println("Changing bp: fp", frontPointer, "\tbp", backPointer, "\tSum", sum, "\tlen", frontPointer-backPointer+1, "\tMax", maxLen)
		}
		if (sum == target) && (frontPointer-backPointer+1 > maxLen) {
			maxLen = frontPointer - backPointer + 1
		}
	}
	// fmt.Println("Max Length : ", maxLen)
	// fmt.Println("---------End--------\n-")
	return maxLen
}
