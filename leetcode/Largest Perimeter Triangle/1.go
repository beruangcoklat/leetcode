func isTriangle(a, b, c int) bool {
	return a+b > c &&
		a+c > b &&
		b+c > a
}

func largestPerimeter(nums []int) int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})
	for i := 0; i < len(nums)-2; i++ {
		if isTriangle(nums[i], nums[i+1], nums[i+2]) {
			return nums[i] + nums[i+1] + nums[i+2]
		}
	}
	return 0
}

