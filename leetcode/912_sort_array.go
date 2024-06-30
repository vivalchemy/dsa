package leetcode

// counting sort
func SortArray(nums []int) []int {
	// -50000 to 50000
	temp := make([]int, 100001)
	for _, value := range nums {
		temp[value+50000]++
	}
	// //debug
	// for i, v := range temp {
	// 	if v == 0 {
	// 		continue
	// 	}
	// 	fmt.Println("index:", i, "value:", v)
	// }
	var indexInNums int
	var indexInTemp int
	for indexInTemp < 100001 {
		if temp[indexInTemp] == 0 {
			indexInTemp++
			continue
		}
		temp[indexInTemp]--
		nums[indexInNums] = indexInTemp - 50000
		indexInNums++
	}

	return nums
}
