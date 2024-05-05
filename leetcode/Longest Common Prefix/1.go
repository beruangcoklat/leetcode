func longestCommonPrefix(strs []string) string {
	idx := 0

	if len(strs) == 1 {
		return strs[0]
	}

	for {
		valid := true
		for i := 0; i < len(strs)-1; i++ {
			if len(strs[i]) <= idx || len(strs[i+1]) <= idx || strs[i][idx] != strs[i+1][idx] {
				valid = false
				break
			}
		}
		if valid {
			idx++
		} else {
			break
		}
	}
	return strs[0][:idx]
}
