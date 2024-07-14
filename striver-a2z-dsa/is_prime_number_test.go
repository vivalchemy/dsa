package strivera2zdsa

import (
	"math"
	"testing"
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

// tests
func TestIsPrimeNumber(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{2, true},   // Smallest prime number
		{3, true},   // Prime number
		{4, false},  // Not a prime number
		{5, true},   // Prime number
		{10, false}, // Not a prime number
		{13, true},  // Prime number
		{1, false},  // Not a prime number
		{0, false},  // Not a prime number
		{-5, false}, // Negative numbers are not prime
	}

	for _, tc := range testCases {
		result := IsPrimeNumber(tc.input)
		if result != tc.expected {
			t.Errorf("For %d, expected %t, but got %t", tc.input, tc.expected, result)
		}
	}
}
