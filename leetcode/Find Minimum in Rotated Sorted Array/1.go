func searchPivotIdx(nums []int, startIdx, endIdx int) int {
	if startIdx > endIdx {
		return -1
	}

	mid := (startIdx + endIdx) / 2

	if mid-1 >= 0 && nums[mid-1] > nums[mid] {
		return mid
	}

	a := searchPivotIdx(nums, startIdx, mid-1)
	if a != -1 {
		return a
	}
	a = searchPivotIdx(nums, mid+1, endIdx)
	if a != -1 {
		return a
	}

	return -1
}

func findMin(nums []int) int {
	pivotIdx := searchPivotIdx(nums, 0, len(nums)-1)
	if pivotIdx == -1 {
		return nums[0]
	}
	if nums[0] < nums[pivotIdx] {
		return nums[0]
	}
	return nums[pivotIdx]
}
