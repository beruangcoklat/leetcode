func minDifference(nums []int) int {
	if len(nums) <= 4 {
		return 0
	}

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	minDiff := 0
	first := true
	for i := 0; i < 4; i++ {
		diff := nums[i+len(nums)-4] - nums[i]
		if first || diff < minDiff {
			first = false
			minDiff = diff
		}
	}
	return minDiff
}