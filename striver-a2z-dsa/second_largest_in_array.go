package strivera2zdsa

func SecondLargestInArray(arr []int) int {
	if len(arr) <= 1 {
		return -1
	}
	max := arr[0]
	max2 := arr[1]
	if max < max2 && len(arr) == 2 {
		return max
	}
	for i := 2; i < len(arr); i++ {
		if arr[i] > max {
			max2 = max
			max = arr[i]
		} else if arr[i] > max2 {
			max2 = arr[i]
		}
	}
	return max2
}
