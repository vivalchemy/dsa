package leetcode

func MinOperations(logs []string) int {
	var count int = 0
	for _, log := range logs {
		switch log {
		case "../":
			if count == 0 {
				continue
			}
			count -= 1
		case "./":
		default:
			count++
		}
	}
	return count
}
