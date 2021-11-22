var memo map[string]int

func lcs(text1, text2 string, idx1, idx2 int) int {
	if idx1 == len(text1) || idx2 == len(text2) {
		return 0
	}

	key := fmt.Sprintf("%v-%v", idx1, idx2)
	if val, ok := memo[key]; ok {
		return val
	}

	ret := 0

	if text1[idx1] == text2[idx2] {
		ret = lcs(text1, text2, idx1+1, idx2+1) + 1
	} else {
		ret = lcs(text1, text2, idx1+1, idx2)
		curr := lcs(text1, text2, idx1, idx2+1)
		if curr > ret {
			ret = curr
		}
	}

	memo[key] = ret
	return ret
}

func minDistance(word1 string, word2 string) int {
	memo = make(map[string]int)
	return len(word1) + len(word2) - (2 * lcs(word1, word2, 0, 0))
}
