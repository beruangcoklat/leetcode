func removeDuplicates(nums []int) int {
	for i := 0; i < len(nums)-1; i++ {
		curr, next := nums[i], nums[i+1]
		if curr == next {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}
