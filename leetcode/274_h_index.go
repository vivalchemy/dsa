package leetcode

import "fmt"

// including the papers that have 1 citation in them
// func HIndex(citations []int) int {
// 	// citations has index at paper and citation of that paper at citations[i]
// 	length := len(citations)
// 	// temp will have citations at i and no of papers having that citations at temp[i]
// 	temp := make([]int, length+1) // 0 to including n
//
// 	// this for loop will convert the citations to the temp array properly
// 	for _, citation := range citations {
// 		if citation > length {
// 			temp[length] += 1
// 		} else {
// 			temp[citation] += 1
// 		}
// 		fmt.Println("citation", citation, "temp: ", temp)
// 	}
//
// 	var totalValidPapersSoFar int
// 	for ; length > 0; length-- {
// 		totalValidPapersSoFar += temp[length]
// 		if totalValidPapersSoFar >= length {
// 			return length
// 		}
// 	}
// 	return 0
// }

// without including the papers that have no citations in them i.e temp[0] will hold papers having 1 citation
func HIndex(citations []int) int {
	length := len(citations)
	temp := make([]int, length) // 1 to including n excluding 0
	for _, citation := range citations {
		if citation == 0 {
			continue
		}
		if citation > length {
			temp[length-1] += 1
		} else {
			temp[citation-1] += 1
		}
		fmt.Println("citation", citation, "temp: ", temp)
	}
	var totalValidPapersSoFar int
	length--
	for ; length >= 0; length-- {
		totalValidPapersSoFar += temp[length]
		if totalValidPapersSoFar > length {
			return length + 1
		}
	}
	return 0
}
