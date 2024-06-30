package strivera2zdsa

import (
	"math"
)

func IsPrimeNumber(num int) bool {
	if num == 2 {
		return true
	} else if num < 2 {
		return false
	}

	sqrt := math.Sqrt(float64(num))

	for i := 2; i <= int(sqrt); i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
