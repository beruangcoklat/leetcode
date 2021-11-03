var memo map[int]int

func solve(idx int, nums []int) int {
	if idx >= len(nums) {
		return 0
	}

	val, ok := memo[idx]
	if ok {
		return val
	}

	take := solve(idx+2, nums) + nums[idx]
	notTake := solve(idx+1, nums)

	max := take
	if take < notTake {
		max = notTake
	}

	memo[idx] = max
	return max
}

func rob(nums []int) int {
	memo = make(map[int]int)
	return solve(0, nums)
}
