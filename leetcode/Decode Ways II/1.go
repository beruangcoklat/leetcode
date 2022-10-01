var memo map[string]int

const mod = 1000000007

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
		mul := 1
		if s[0] == '*' {
			mul = 9
		}
		ret += solve(s[1:]) * mul
	}

	if length >= 2 {
		if s[0] != '*' && s[1] != '*' {
			if s[0] == '1' || (s[0] == '2' && s[1] <= '6') {
				ret += solve(s[2:])
			}
		} else if s[0] == '*' && s[1] != '*' {
			mul := 1
			if s[1] <= '6' {
				mul = 2
			}
			ret += solve(s[2:]) * mul
		} else if (s[0] == '1' || s[0] == '2') && s[1] == '*' {
			mul := 1
			if s[0] == '1' {
				mul = 9
			} else if s[0] == '2' {
				mul = 6
			}
			ret += solve(s[2:]) * mul
		} else if s[0] == '*' && s[1] == '*' {
			ret += solve(s[2:]) * 15
		}
	}

	ret %= mod
	memo[s] = ret
	return ret
}

func numDecodings(s string) int {
	memo = make(map[string]int)
	return solve(s)
}
