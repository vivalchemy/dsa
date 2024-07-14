package strivera2zdsa

func getSingleElement(nums []int) int {
	var result int
	for _, value := range nums {
		result ^= value
	}
	return result
}
