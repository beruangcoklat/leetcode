var cache map[int]int

func solve(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 || n == 2 {
		return 1
	}
	val, ok := cache[n]
	if ok {
		return val
	}
	ret := solve(n-1) + solve(n-2) + solve(n-3)
	cache[n] = ret
	return ret
}

func tribonacci(n int) int {
	cache = make(map[int]int)
	return solve(n)
}
