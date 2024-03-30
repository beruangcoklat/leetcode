func isIsomorphic(s string, t string) bool {
	size := len(s)
	sMap := make(map[byte]int)
	tMap := make(map[byte]int)
	for i := 0; i < size; i++ {
		sInt, sOk := sMap[s[i]]
		tInt, tOk := tMap[t[i]]
		if (sOk && !tOk) || (!sOk && tOk) || (sInt != tInt) {
			return false
		}
		sMap[s[i]] = i
		tMap[t[i]] = i
	}
	return true
}
