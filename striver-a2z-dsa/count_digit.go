package strivera2zdsa

import "math"

func CountDigits(n int) int {
	return int(math.Log10(float64(n)) + 1)
}
