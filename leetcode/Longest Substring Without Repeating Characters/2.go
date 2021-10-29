func isRepeating(str string) bool {
	strMap := make(map[byte]struct{})
	for i := 0; i < len(str); i++ {
		if _, ok := strMap[str[i]]; ok {
			return true
		}
		strMap[str[i]] = struct{}{}
	}
	return false
}

func lengthOfLongestSubstring(s string) int {
	length := len(s)

	if !isRepeating(s) {
		return length
	}

	ret := 1
	currLength := 2
	for currLength < length {
		allRepeating := true
		for i := 0; i+currLength <= length; i++ {
			if !isRepeating(s[i : i+currLength]) {
				allRepeating = false
				break
			}
		}

		if allRepeating {
			break
		}

		if currLength > ret {
			ret = currLength
		}

		currLength++
	}

	return ret
}
