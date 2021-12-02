func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(nums []int, startIdx int) int {
	curr := 1
	max := 0
	first := true
	for i := startIdx; i < len(nums); i++ {
		curr *= nums[i]
		if first || curr > max {
			max = curr
		}
		first = false
	}
	return max
}

func maxProduct(nums []int) int {
	max := 0
	first := true
	for i := 0; i < len(nums); i++ {
		curr := solve(nums, i)
		if first || curr > max {
			max = curr
		}
		first = false
	}
	return max
}
