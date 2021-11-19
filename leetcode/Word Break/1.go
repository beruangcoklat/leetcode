var memo map[string]struct{}

func solve(s string, wordDict []string) bool {
	if s == "" {
		return true
	}

	if _, ok := memo[s]; ok {
		return false
	}

	for _, word := range wordDict {
		if len(word) > len(s) {
			continue
		}

		if s[:len(word)] != word {
			continue
		}

		if solve(s[len(word):], wordDict) {
			return true
		}
	}

	memo[s] = struct{}{}
	return false
}

func wordBreak(s string, wordDict []string) bool {
	memo = make(map[string]struct{})
	return solve(s, wordDict)
}
