package leetcode

// hashmap
func TwoSum(nums []int, target int) []int {
	bal := make(map[int]int, len(nums))
	result := make([]int, 0, 2)
	for index, num := range nums {
		companion, exist := bal[num]
		if exist {
			result = append(result, index, companion)
			return result
		}
		bal[target-num] = index
	}
	return nil
}
