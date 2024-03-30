func isIsomorphic(sArr []string, tArr []string) bool {
	size := len(sArr)

	if size != len(tArr) {
		return false
	}

	sMap := make(map[string]int)
	tMap := make(map[string]int)
	for i := 0; i < size; i++ {
		sInt, sOk := sMap[sArr[i]]
		tInt, tOk := tMap[tArr[i]]
		if (sOk && !tOk) || (!sOk && tOk) || (sInt != tInt) {
			return false
		}
		sMap[sArr[i]] = i
		tMap[tArr[i]] = i
	}
	return true
}

func wordPattern(pattern string, s string) bool {
	tokens := strings.Split(s, " ")
	patternArr := make([]string, len(pattern))
	for i, c := range pattern {
		patternArr[i] = string(c)
	}
	return isIsomorphic(patternArr, tokens)
}