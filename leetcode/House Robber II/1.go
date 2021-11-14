func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(idx int, nums []int, canRobLastHouse bool) int {
	if idx >= len(nums) {
		return 0
	}

	key := fmt.Sprintf("%v-%v", idx, canRobLastHouse)
	if val, ok := memo[key]; ok {
		return val
	}

	notTake := solve(idx+1, nums, canRobLastHouse)
	max := notTake

	if idx == 0 {
		take := solve(idx+2, nums, false) + nums[idx]
		max = getMax(notTake, take)
	} else if (idx == len(nums)-1 && canRobLastHouse) || (idx < len(nums)-1) {
		take := solve(idx+2, nums, canRobLastHouse) + nums[idx]
		max = getMax(notTake, take)
	}

	memo[key] = max
	return max
}

var memo map[string]int

func rob(nums []int) int {
	memo = make(map[string]int)
	return solve(0, nums, true)
}