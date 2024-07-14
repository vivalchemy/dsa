package strivera2zdsa

// https://takeuforward.org/arrays/find-the-missing-number-in-an-array/
func FindMissingNumber(nums []int) int {
	var result int = len(nums) + 1
	for index, value := range nums {
		result ^= value ^ (index + 1)
	}
	return result
}
