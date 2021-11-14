var memo map[int]struct{}

func solve(idx int, nums []int) bool {
	if idx == len(nums)-1 {
		return true
	}

	if _, ok := memo[idx]; ok {
		return false
	}

	for i := 1; i <= nums[idx]; i++ {
		if idx+i > len(nums) {
			continue
		}
		if solve(idx+i, nums) {
			return true
		}
	}

	memo[idx] = struct{}{}
	return false
}

func canJump(nums []int) bool {
	memo = make(map[int]struct{})
	return solve(0, nums)
}