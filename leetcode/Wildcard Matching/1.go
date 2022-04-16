func substring(s string, lenS, idx int) string {
	if idx >= lenS {
		return ""
	}
	return s[idx : idx+1]
}

var memo map[string]bool = make(map[string]bool)

func solve(s, p string, idxS, idxP int) bool {
	key := fmt.Sprintf("%v-%v", idxS, idxP)
	if val, ok := memo[key]; ok {
		return val
	}

	lenS := len(s)
	lenP := len(p)

	if idxP == lenP {
		return idxS == lenS
	}

	if idxS == lenS {
		res := true
		for i := idxP; i < lenP; i++ {
			if p[i] != '*' {
				res = false
				break
			}
		}
		memo[key] = res
		return res
	}

	s1 := substring(s, lenS, idxS)
	p1 := substring(p, lenP, idxP)
	ret := false

	if s1 == p1 && solve(s, p, idxS+1, idxP+1) {
		ret = true
	} else if p1 == "?" && solve(s, p, idxS+1, idxP+1) {
		ret = true
	} else if p1 == "*" && (solve(s, p, idxS+1, idxP+1) || solve(s, p, idxS+1, idxP) || solve(s, p, idxS, idxP+1)) {
		ret = true
	}

	memo[key] = ret
	return ret
}

func isMatch(s string, p string) bool {
	memo = make(map[string]bool)
	return solve(s, p, 0, 0)
}
