package strivera2zdsa

func LongestSubarrayWithGivenSumK(nums []int, target int) int {
	var maxLength int
	var sum int
	sumMap := make(map[int]int)
	sumMap[0] = -1
	for index, num := range nums {
		sum += num
		if startIndex, exists := sumMap[sum-target]; (exists) && (maxLength < index-startIndex) {
			maxLength = index - startIndex
		}
		if _, exists := sumMap[sum]; !exists {
			sumMap[sum] = index
		}
	}
	return maxLength
}
