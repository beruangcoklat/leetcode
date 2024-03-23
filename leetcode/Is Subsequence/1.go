func isSubsequence(s string, t string) bool {
	sizeS := len(s)
	sizeT := len(t)
	idxS := 0
	idxT := 0

	if sizeS == 0 {
		return true
	}

	for idxS < sizeS && idxT < sizeT {
		if s[idxS] == t[idxT] {
			idxS++
			idxT++
		} else {
			idxT++
		}
	}

	return idxS == sizeS
}
