package strivera2zdsa

import "testing"

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

// tests
func TestGCD(t *testing.T) {
	testCases := []struct {
		num1     int
		num2     int
		expected int
	}{
		{10, 25, 5},
		{14, 28, 14},
		{21, 7, 7},
		{35, 15, 5},
		{81, 27, 27},
	}

	for _, tc := range testCases {
		got := GCD(tc.num1, tc.num2)
		if got != tc.expected {
			t.Errorf("GCD(%d, %d) = %d; want %d", tc.num1, tc.num2, got, tc.expected)
		}
	}
}
