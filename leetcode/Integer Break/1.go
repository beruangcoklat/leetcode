var memo map[int]int

func solve(n int, depth int) int {
	if n == 0 {
		if depth < 2 {
			return 0
		}
		return 1
	}

	val, ok := memo[n]
	if ok {
		return val
	}

	first := true
	max := 0
	for i := 1; i <= n; i++ {
		curr := solve(n-i, depth+1) * i
		if first || curr > max {
			max = curr
			first = false
		}
	}

	memo[n] = max
	return max
}

func integerBreak(n int) int {
	memo = make(map[int]int)
	return solve(n, 0)
}
