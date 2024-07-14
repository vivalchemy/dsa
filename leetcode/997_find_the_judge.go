package leetcode

// keeps the track of both the from and to of trust.
// func FindJudge(n int, trust [][]int) int {
// 	hasTrust := make([][2]int, n+1)
// 	for _, t := range trust {
// 		hasTrust[t[1]][1]++
// 		hasTrust[t[0]][0]++
// 	}
// 	for index := 1; index <= n; index++ {
// 		if hasTrust[index][0] == 0 && hasTrust[index][1] == n-1 {
// 			return index
// 		}
// 	}
//
// 	return -1
// }

// if a person trusts someone then put -1 in his place
func FindJudge(n int, trust [][]int) int {
	/*
	 need to explicity mention this condition since this for this condition
	 the first for loop is skipped and the in the second for loop both the
	 index 0 and 1 satisfy the condition so instead of 1 it will return 0
	*/
	if n == 1 {
		return 1
	}
	hasTrust := make([]int, n+1)
	for _, t := range trust {
		hasTrust[t[0]] = -1
		if hasTrust[t[1]] != -1 {
			hasTrust[t[1]]++
		}
	}
	for index, trustedBy := range hasTrust {
		// since index 0 will not be n - 1 it won't matter whether we include it or not
		if trustedBy == n-1 {
			return index
		}
	}
	return -1
}
