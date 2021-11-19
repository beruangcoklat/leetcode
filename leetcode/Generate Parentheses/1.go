var ret []string

func solve(open, close int, n int, res string) {
	if open == close && open == n {
		ret = append(ret, res)
		return
	}
	if open < n {
		solve(open+1, close, n, res+"(")
	}
	if close < open {
		solve(open, close+1, n, res+")")
	}
}

func generateParenthesis(n int) []string {
	ret = []string{}
	solve(0, 0, n, "")
	return ret
}
