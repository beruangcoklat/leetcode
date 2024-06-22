func longestPalindrome(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}

	ret := 0
	oddFound := 0
	for _, count := range m {
		if count%2 == 0 {
			ret += count
		} else {
			ret += count - 1
			oddFound = 1
		}
	}
	return ret + oddFound
}