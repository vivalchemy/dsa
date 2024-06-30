package leetcode

func RomanToInt(s string) int {
	var result int
	m := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	sLen := len(s)
	for i, digit := range s {
		if i+1 < sLen && m[digit] < m[rune(s[i+1])] {
			result -= m[digit]
		} else {
			result += m[digit]
		}
	}
	return result
}
