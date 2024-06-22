func appendCharacters(s string, t string) int {
	sIdx := 0
	tIdx := 0

	for sIdx < len(s) && tIdx < len(t) {
		if t[tIdx] == s[sIdx] {
			sIdx++
			tIdx++
		} else {
			sIdx++
		}
	}

	return len(t) - tIdx
}

