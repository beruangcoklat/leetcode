var cache [301]int

func solve(idx int, s string, wordDict []string) bool {
	if idx == len(s) {
		return true
	}

	if cache[idx] != 0 {
		return cache[idx] == 2
	}

	for _, word := range wordDict {
		if idx+len(word) > len(s) {
			continue
		}
		if s[idx:idx+len(word)] == word && solve(idx+len(word), s, wordDict) {
			cache[idx] = 2
			return true
		}
	}

	cache[idx] = 1
	return false
}

func wordBreak(s string, wordDict []string) bool {
	cache = [301]int{}
	return solve(0, s, wordDict)
}
