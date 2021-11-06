func search(nums []int, startIdx, endIdx int) int {
	if startIdx > endIdx {
		return -1
	}

	mid := (startIdx + endIdx) / 2

	if mid-1 >= 0 && mid+1 <= len(nums)-1 && nums[mid-1] < nums[mid] && nums[mid+1] < nums[mid] {
		return mid
	}

	a := search(nums, startIdx, mid-1)
	if a != -1 {
		return a
	}
	a = search(nums, mid+1, endIdx)
	if a != -1 {
		return a
	}

	return -1
}

func findPeakElement(nums []int) int {
	idx := search(nums, 0, len(nums)-1)
	if idx != -1 {
		return idx
	}

	max := 0
	for i := 1; i < len(nums); i++ {
		if nums[i] > nums[max] {
			max = i
		}
	}
	return max
}

