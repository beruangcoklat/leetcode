func lengthOfLastWord(s string) int {
	startIdx := -1
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' && startIdx == -1 {
			continue
		}

		if s[i] == ' ' {
			return startIdx - i
		}

		if startIdx == -1 {
			startIdx = i
			continue
		}
	}

	if startIdx == -1 {
		return 0
	}
	return startIdx + 1
}
