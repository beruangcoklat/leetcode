func removeDuplicates(nums []int) int {
	if len(nums) < 3 {
		return len(nums)
	}
	clone := make([]int, len(nums))
	copy(clone, nums)
	idx := 2
	for i := 2; i < len(nums); i++ {
		if clone[i] == clone[i-1] && clone[i-1] == clone[i-2] {
			continue
		}
		nums[idx] = nums[i]
		idx++
	}
	return idx
}