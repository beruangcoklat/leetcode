var globalResult int

func subsetXORSum(nums []int) int {
	globalResult = 0
	solve(-1, nums, make([]int, 0))
	return globalResult
}

func solve(idx int, nums, result []int) {
	if len(result) > 0 {
		globalResult += xor(result)
	}
	if len(result) == len(nums) {
		return
	}
	for i := idx + 1; i < len(nums); i++ {
		solve(i, nums, clone(result, nums[i]))
	}
}

func clone(nums []int, num int) []int {
	result := make([]int, len(nums)+1)
	copy(result, nums)
	result[len(nums)] = num
	return result
}

func xor(nums []int) int {
	result := nums[0]
	for i := 1; i < len(nums); i++ {
		result ^= nums[i]
	}
	return result
}
