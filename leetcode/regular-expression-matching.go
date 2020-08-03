// https://leetcode.com/problems/regular-expression-matching/

func get(s string, lenS, idx int) string {
	if idx >= lenS {
		return ""
	}
	return s[idx : idx+1]
}

var memo map[string]bool = make(map[string]bool)

func isMatch(s string, p string) bool {
	if val, ok := memo[s+"###"+p]; ok {
		return val
	}

	lenS := len(s)
	lenP := len(p)

	if lenP == 0 {
		return lenS == 0
	}

	ret := false

	s1 := get(s, lenS, 0)
	p1 := get(p, lenP, 0)
	p2 := get(p, lenP, 1)

	// match
	if (s1 == p1 || p1 == ".") && lenS > 0 {
		// next both
		if isMatch(s[1:], p[1:]) {
			ret = true
		}

		// next s
		if p2 == "*" && isMatch(s[1:], p) {
			ret = true
		}

		// * finish
		if p2 == "*" && isMatch(s[1:], p[2:]) {
			ret = true
		}
	}

	// can be removed
	if p2 == "*" && isMatch(s, p[2:]) {
		ret = true
	}

	memo[s+"###"+p] = ret
	return ret
}
