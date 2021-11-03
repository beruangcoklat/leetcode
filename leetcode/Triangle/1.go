var memo map[int]map[int]int

func accessMemo(a, b int) (int, bool) {
	if memo[a] == nil {
		return 0, false
	}
	memoMap, ok := memo[a]
	if !ok {
		return 0, false
	}
	val, ok := memoMap[b]
	if !ok {
		return 0, false
	}
	return val, true
}

func storeMemo(a, b, val int) {
	if memo[a] == nil {
		memo[a] = make(map[int]int)
	}
	memo[a][b] = val
}

func solve(height, idx int, triangle [][]int) int {
	if height == len(triangle) {
		return 0
	}

	val, ok := accessMemo(height, idx)
	if ok {
		return val
	}

	a := triangle[height][idx] + solve(height+1, idx, triangle)
	b := triangle[height][idx] + solve(height+1, idx+1, triangle)
	res := a
	if b < res {
		res = b
	}

	storeMemo(height, idx, res)
	return res
}

func minimumTotal(triangle [][]int) int {
	memo = make(map[int]map[int]int)
	return solve(0, 0, triangle)
}
