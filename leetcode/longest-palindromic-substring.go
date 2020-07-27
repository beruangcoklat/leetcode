// https://leetcode.com/problems/longest-palindromic-substring/

func isPalindrome(a string) bool {
	strLen := len(a)
	if strLen == 0 {
		return false
	}
	for i := 0; i < strLen/2; i++ {
		if a[i] != a[strLen-i-1] {
			return false
		}
	}
	return true
}

func longestPalindrome(s string) string {
	if len(s) == 0 {
		return s
	}
	max := s[0:1]
	maxLen := 1
	lenS := len(s)
	for i := 0; i < lenS; i++ {
		for j := 2; j < lenS-i+1; j++ {
			c := s[i:(i + j)]
			if c == max {
				continue
			}
			if lenC := len(c); lenC > maxLen && isPalindrome(c) {
				max = c
				maxLen = lenC
			}
		}
	}
	return max
}
