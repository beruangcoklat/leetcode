func minWindow(s string, t string) string {
	l, r := 0, 0
	charMap := make(map[string]int)
	charMap[string(s[0])]++

	tMap := make(map[string]int)
	for _, char := range t {
		tMap[string(char)]++
	}

	minStr := ""
	for l <= r && r < len(s) {
		length := r - l + 1
		if valid(charMap, tMap) {
			if minStr == "" || length < len(minStr) {
				minStr = s[l : r+1]
			}
			charMap[string(s[l])]--
			l++
		} else {
			r++
			if r == len(s) {
				break
			}
			charMap[string(s[r])]++
		}
	}
	return minStr
}

func valid(charMap map[string]int, tMap map[string]int) bool {
	for k, v := range tMap {
		if charMap[k] < v {
			return false
		}
	}
	return true
}
