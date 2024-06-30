func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}
	l, r := 0, 0
	max := 0
	charMap := make(map[string]int)
	charMap[string(s[r])]++

	for l <= r && r < len(s) {
		length := r - l + 1

		if !isRepeating(charMap) {
			if length > max {
				max = length
			}
			r++
			if r == len(s) {
				break
			}
			charMap[string(s[r])]++
		} else {
			charMap[string(s[l])]--
			l++
		}
	}
	return max
}

func isRepeating(charMap map[string]int) bool {
	for _, count := range charMap {
		if count > 1 {
			return true
		}
	}
	return false
}
