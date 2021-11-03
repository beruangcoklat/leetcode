var (
	memo map[int]int
)

func climbStairs(n int) int {
	if n == 0 {
		return 1
	}
	if n < 0 {
		return 0
	}

	if memo == nil {
		memo = make(map[int]int)
	}
	val, ok := memo[n]
	if ok {
		return val
	}

	res := climbStairs(n-1) + climbStairs(n-2)
	memo[n] = res
	return res
}
