var memo map[string]int

func solve(height, idx int, triangle [][]int) int {
	if height == len(triangle) {
		return 0
	}

	key := fmt.Sprintf("%v-%v", height, idx)
	val, ok := memo[key]
	if ok {
		return val
	}

	a := triangle[height][idx] + solve(height+1, idx, triangle)
	b := triangle[height][idx] + solve(height+1, idx+1, triangle)
	res := a
	if b < res {
		res = b
	}

	memo[key] = res
	return res
}

func minimumTotal(triangle [][]int) int {
	memo = make(map[string]int)
	return solve(0, 0, triangle)
}
