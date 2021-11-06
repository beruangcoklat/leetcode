func searchMatrix(matrix [][]int, target int) bool {
	nums := []int{}
	for _, row := range matrix {
		nums = append(nums, row...)
	}

	l, r := 0, len(nums)-1

	for l <= r {
		mid := (l + r) / 2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return false
}
