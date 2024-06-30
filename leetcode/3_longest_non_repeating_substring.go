package leetcode

// hashmao
func LengthOfLongestSubstring(s string) int {
	var maxLength int
	var length int
	m := map[rune]int{}
	for i, char := range s {
		if val, exists := m[char]; exists {
			if length > maxLength {
				maxLength = length
			}
			length = i - val
			for k, v := range m {
				if v <= val {
					delete(m, k)
				}
			}
		} else {
			length += 1
		}
		m[char] = i
	}
	if length > maxLength {
		maxLength = length
	}
	return maxLength
}

// sliding window
func initSliceToAValue(length, value int) []int {
	slice := make([]int, length)
	if length > 0 {
		slice[0] = value
		for i := 1; i < length; i *= 2 {
			copy(slice[i:], slice[:i])
		}
	}
	return slice
}

func LengthOfLongestSubstringAlternateUsingArray(s string) int {
	var maxLength int
	var charArr = initSliceToAValue(128, -1)
	var left, right int
	sLen := len(s)
	for right = 0; right < sLen; right++ {
		// fmt.Println("\nleft", left, "\nright", right, "\nmaxLength", maxLength, "\ns[right]", s[right], "\ncharArr[s[right]]", charArr[s[right]])
		if charArr[s[right]] >= left {
			maxLength = max(maxLength, right-left)
			left = charArr[s[right]] + 1
		}
		charArr[s[right]] = right
	}
	maxLength = max(maxLength, right-left)
	return maxLength
}
