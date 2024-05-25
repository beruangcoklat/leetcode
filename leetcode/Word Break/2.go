var cache map[int]bool

func solve(idx int, s string, wordDict []string) bool {
	if idx == len(s) {
		return true
	}

	val, ok := cache[idx]
	if ok {
		return val
	}

	for _, word := range wordDict {
		if idx+len(word) > len(s) {
			continue
		}
		if s[idx:idx+len(word)] == word && solve(idx+len(word), s, wordDict) {
			cache[idx] = true
			return true
		}
	}

	cache[idx] = false
	return false
}

func wordBreak(s string, wordDict []string) bool {
	cache = make(map[int]bool)
	return solve(0, s, wordDict)
}