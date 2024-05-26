var cache [100001][2][3]int

func checkRecord(n int) int {
	cache = [100001][2][3]int{}
	return solve(n, 0, 0, 0)
}

func solve(n int, lastLateCount int, absentCount int, length int) int {
	if length == n {
		return 1
	}

	val := cache[length][absentCount][lastLateCount]
	if val != 0 {
		return val
	}

	ret := 0

	// late
	if lastLateCount < 2 {
		ret += solve(n, lastLateCount+1, absentCount, length+1)
		ret %= 1000000007
	}

	// absent
	if absentCount < 1 {
		ret += solve(n, 0, absentCount+1, length+1)
		ret %= 1000000007
	}

	// present
	ret += solve(n, 0, absentCount, length+1)
	ret %= 1000000007

	cache[length][absentCount][lastLateCount] = ret
	return ret
}
