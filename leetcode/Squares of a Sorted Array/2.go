func sortedSquares(nums []int) []int {
	l, r := 0, len(nums)-1
	ret := make([]int, len(nums))
	idx := len(nums) - 1

	for l <= r {
		a := nums[l] * nums[l]
		b := nums[r] * nums[r]
		if a >= b {
			l++
			ret[idx] = a
		} else {
			r--
			ret[idx] = b
		}
		idx--
	}

	return ret
}
