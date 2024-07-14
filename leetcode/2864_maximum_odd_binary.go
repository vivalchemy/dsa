package leetcode

func maximumOddBinaryNumber(s string) string {
	n := len(s)
	if n <= 1 {
		return s
	}
	var zeroCounter int
	soln := make([]rune, n)
	for _, char := range s {
		if char == '0' {
			zeroCounter++
		}
	}
	for i := 0; i < n-zeroCounter-1; i++ {
		soln[i] = '1'
	}
	for i := n - zeroCounter - 1; i < n-1; i++ {
		soln[i] = '0'
	}
	soln[n-1] = '1'
	return string(soln)
}
