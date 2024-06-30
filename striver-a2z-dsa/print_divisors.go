package strivera2zdsa

import (
	"math"
	"sort"
)

func PrintDivisor(n int) []int {
	sqrt := math.Sqrt(float64(n))
	soln := make([]int, 0)
	for i := 1; i <= int(sqrt); i++ {
		if n%i == 0 {
			soln = append(soln, i)
			if n/i != i {
				soln = append(soln, n/i)
			}
		}
	}
	sort.Ints(soln)
	return soln
}
