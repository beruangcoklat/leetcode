func minDifference(nums []int) int {
	if len(nums) <= 3 {
		return 0
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	l, r := 0, len(nums)-1
	return solve(nums, l, r, 0)
}

func solve(nums []int, l int, r int, move int) int {
	if move == 3 {
		return nums[r] - nums[l]
	}
	a := solve(nums, l+1, r, move+1)
	b := solve(nums, l, r-1, move+1)
	if a < b {
		return a
	}
	return b
}