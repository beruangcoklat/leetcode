var globalResult []string

func solve(idx int, s string, wordDict []string, result string) {
	if idx == len(s) {
		globalResult = append(globalResult, result)
		return
	}

	for _, word := range wordDict {
		if idx+len(word) > len(s) {
			continue
		}
		if s[idx:idx+len(word)] == word {
			var newResult string
			if result == "" {
				newResult = word
			} else {
				newResult = result + " " + word
			}

			solve(idx+len(word), s, wordDict, newResult)
		}
	}
}

func wordBreak(s string, wordDict []string) []string {
	globalResult = make([]string, 0)
	solve(0, s, wordDict, "")
	return globalResult
}
