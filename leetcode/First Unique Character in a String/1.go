func firstUniqChar(s string) int {
	m := make(map[rune]int)
	for _, c := range s {
		m[c]++
	}
	for i := 0; i < len(s); i++ {
		if m[rune(s[i])] == 1 {
			return i
		}
	}
	return -1
}
