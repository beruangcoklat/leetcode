func substring(s string, lenS, idx int) string {
	if idx >= lenS {
		return ""
	}
	return s[idx : idx+1]
}

var memo map[string]bool

func solve(s, p string, idxS, idxP int) bool {
	key := fmt.Sprintf("%v-%v", idxS, idxP)
	if val, ok := memo[key]; ok {
		return val
	}

	ret := false
	defer func() {
		memo[key] = ret
	}()

	lenS := len(s)
	lenP := len(p)

	if idxP == lenP {
		return idxS == lenS
	}

	if idxS == lenS {
		ret = true
		for i := idxP; i < lenP; i++ {
			if p[i] != '*' && (i+1 >= lenP || p[i+1] != '*') {
				ret = false
				return ret
			}
		}
		return ret
	}

	s1 := substring(s, lenS, idxS)
	p1 := substring(p, lenP, idxP)
	p2 := substring(p, lenP, idxP+1)

	if (s1 == p1 || p1 == ".") && solve(s, p, idxS+1, idxP+1) {
		ret = true
		return ret
	}

	if (s1 == p1 || p1 == ".") && p2 == "*" && (solve(s, p, idxS+1, idxP) || solve(s, p, idxS+1, idxP+2)) {
		ret = true
		return ret
	}

	if p2 == "*" && solve(s, p, idxS, idxP+2) {
		ret = true
		return ret
	}

	return ret
}

func isMatch(s string, p string) bool {
	memo = make(map[string]bool)
	return solve(s, p, 0, 0)
}
