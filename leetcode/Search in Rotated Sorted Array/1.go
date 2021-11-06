func searchPivotIdx(nums []int, target, startIdx, endIdx int) int {
	if startIdx > endIdx {
		return -1
	}

	mid := (startIdx + endIdx) / 2

	if mid-1 >= 0 && nums[mid-1] > nums[mid] {
		return mid
	}

	a := searchPivotIdx(nums, target, startIdx, mid-1)
	if a != -1 {
		return a
	}
	a = searchPivotIdx(nums, target, mid+1, endIdx)
	if a != -1 {
		return a
	}

	return -1
}

func search(nums []int, target int) int {
	pivotIdx := searchPivotIdx(nums, target, 0, len(nums)-1)

	binarySearch := func(l, r int) int {
		for l <= r {
			mid := (l + r) / 2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] < target {
				l = mid + 1
			} else {
				r = mid - 1
			}
		}
		return -1
	}

	if pivotIdx == -1 {
		return binarySearch(0, len(nums)-1)
	}

	a := binarySearch(0, pivotIdx-1)
	if a != -1 {
		return a
	}

	a = binarySearch(pivotIdx, len(nums)-1)
	if a != -1 {
		return a
	}

	return -1
}
