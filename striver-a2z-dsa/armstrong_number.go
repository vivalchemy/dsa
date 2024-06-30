package strivera2zdsa

import "math"

func ArmstrongNumber(n int) bool {
	numberOfDigits := int(math.Log10(float64(n)) + 1)
	temp := n
	var raisedNum int
	for temp > 0 {
		raisedNum += int(math.Pow(float64(temp%10), float64(numberOfDigits)))
		temp /= 10
	}
	return raisedNum == n
}
