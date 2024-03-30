func solve(s string, t string) bool {
	size := len(s)

	m := make(map[byte]byte)
	for i := 0; i < size; i++ {
		if val, ok := m[s[i]]; ok {
			if val != t[i] {
				return false
			}
		}
		m[s[i]] = t[i]
	}

	return true
}

func isIsomorphic(s string, t string) bool {
	return solve(s, t) && solve(t, s)
}
