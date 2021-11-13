func isAnagram(wordA, wordB map[byte]int) bool {
	for k, v := range wordA {
		if wordB[k] != v {
			return false
		}
	}
	return true
}

func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return []int{}
	}

	wordP := make(map[byte]int)
	wordS := make(map[byte]int)

	for i := 0; i < len(p); i++ { // _, s := range p {
		wordP[p[i]]++
		wordS[s[i]]++
	}

	ret := []int{}
	if isAnagram(wordP, wordS) {
		ret = append(ret, 0)
	}

	for i := len(p); i < len(s); i++ {
		wordS[s[i-len(p)]]--
		wordS[s[i]]++
		if isAnagram(wordP, wordS) {
			ret = append(ret, i-len(p)+1)
		}
	}

	return ret
}
