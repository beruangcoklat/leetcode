func rotate(nums []int, k int) {
	ret := make([]int, len(nums))
	for i := range nums {
		idx := i - (k % len(nums))
		if idx < 0 {
			idx = len(nums) + idx
		}
		ret[i] = nums[idx]
	}
	for i, v := range ret {
		nums[i] = v
	}
}
