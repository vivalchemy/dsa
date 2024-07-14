package strivera2zdsa

func GetSingleElement(nums []int) int {
	var result int
	for _, value := range nums {
		result ^= value
	}
	return result
}
