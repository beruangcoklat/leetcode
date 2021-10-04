func searchInsert(nums []int, target int) int {
	left, right := 0, len(nums)-1
	mid := 0
	for left <= right {
		mid = (left + right) / 2
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	if target > nums[mid] {
		return mid + 1
	}

	return mid
}
