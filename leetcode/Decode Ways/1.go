var memo map[string]int

func solve(s string) int {
	length := len(s)
	if length == 0 {
		return 1
	}

	if mem, ok := memo[s]; ok {
		return mem
	}

	ret := 0
	if s[0] != '0' {
		ret += solve(s[1:])
	}

	if length >= 2 {
		if s[0] == '1' || (s[0] == '2' && s[1] <= '6') {
			ret += solve(s[2:])
		}
	}

	memo[s] = ret
	return ret
}

func numDecodings(s string) int {
	memo = make(map[string]int)
	return solve(s)
}
