package leetcode

import (
	"fmt"
	"math"
)

func ReverseInt(x int) int {
	var rev int
	// true pos false neg
	var sign bool
	if x >= 0 {
		sign = true
	}
	if !sign {
		x *= -1
	}
	for x > 0 {
		rev = rev*10 + x%10
		if rev > math.MaxInt32 || (!sign && rev > math.MaxInt32+1) {
			return 0
		}
		x /= 10
		fmt.Println("\trev:", rev, "\tx:", x)
	}
	if !sign {
		return rev * -1
	}
	return rev
}
