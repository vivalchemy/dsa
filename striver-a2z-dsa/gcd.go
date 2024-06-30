package strivera2zdsa

func GCD(num1, num2 int) int {
	if num1 < 0 {
		num1 = -num1
	} else if num2 < 0 {
		num2 = -num2
	}
	if num1 == 0 || num1%num2 == 0 {
		return num2
	} else if num2 == 0 || num2%num1 == 0 {
		return num1
	}
	for num2 > 0 {
		num1, num2 = num2, num1%num2
	}
	return num1
}
