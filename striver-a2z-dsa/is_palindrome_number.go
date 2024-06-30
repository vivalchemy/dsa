package strivera2zdsa

func IsPalindromeNumber(n int) bool {
	// basic check before randoms
	if n != 0 && (n < 0 || n%10 == 0) {
		return false
	} else if n < 10 {
		return true
	}
	var rev int
	temp := n
	for temp > 0 {
		rev = rev*10 + temp%10
		temp /= 10
	}
	return rev == n
}
