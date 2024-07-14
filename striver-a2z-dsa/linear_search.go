package strivera2zdsa

func LinearSearch(nums []int, target int) int {
	for index, value := range nums {
		if value == target {
			return index
		}
	}
	return -1
}
