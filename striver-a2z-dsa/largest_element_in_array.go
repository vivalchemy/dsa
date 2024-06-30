package strivera2zdsa

func LargestElementInArray(arr []int) int {
	if len(arr) == 0 {
		return 0
	}
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}
