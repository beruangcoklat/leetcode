func lengthOfLongestSubstring(s string) int {
	size := len(s)
	if size == 0 {
		return 0
	}

	charCountMap := make(map[string]int)

	leftIdx := 0
	rightIdx := 0
	charCountMap[string(s[leftIdx])]++
	max := 1

	for {
		rightIdx++
		if rightIdx >= size {
			break
		}

		charCountMap[string(s[rightIdx])]++
		for charCountMap[string(s[rightIdx])] > 1 {
			charCountMap[string(s[leftIdx])]--
			leftIdx++
		}

		length := rightIdx - leftIdx + 1
		if length > max {
			max = length
		}
	}

	return max
}
