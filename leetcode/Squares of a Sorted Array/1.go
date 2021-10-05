func sortedSquares(nums []int) []int {
	for i := range nums {
		nums[i] = nums[i] * nums[i]
	}
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	return nums
}
