func searchLeft(nums []int, target, startIdx, endIdx int) int {
	if nums[startIdx] == target {
		return startIdx
	}

	l, r := startIdx, endIdx
	ret := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
			continue
		}
		if nums[mid] > target {
			r = mid - 1
			continue
		}
		ret = mid
		if mid-1 < startIdx || nums[mid-1] != nums[mid] {
			break
		} else {
			r = mid - 1
		}
	}
	return ret
}

func searchRight(nums []int, target, startIdx, endIdx int) int {
	if nums[endIdx] == target {
		return endIdx
	}
	l, r := startIdx, endIdx
	ret := -1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
			continue
		}
		if nums[mid] > target {
			r = mid - 1
			continue
		}
		ret = mid
		if mid+1 > endIdx || nums[mid+1] != nums[mid] {
			break
		} else {
			l = mid + 1
		}
	}
	return ret
}

func searchRange(nums []int, target int) []int {
	l, r := 0, len(nums)-1
	for l <= r {
		mid := (l + r) / 2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			left := searchLeft(nums, target, 0, mid)
			right := searchRight(nums, target, mid, len(nums)-1)
			if left == -1 && right != -1 {
				left = right
			} else if right == -1 && left != -1 {
				right = left
			}
			return []int{left, right}
		}
	}
	return []int{-1, -1}
}

