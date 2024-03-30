func countSubarrays(nums []int, k int) int64 {
	size := len(nums)
	if size == 0 {
		return 0
	}

	maxNum := nums[0]
	for i := 1; i < size; i++ {
		if nums[i] > maxNum {
			maxNum = nums[i]
		}
	}

	var result int64

	startIdx := 0
	endIdx := 0
	count := 0

	for endIdx < size {
		if nums[endIdx] == maxNum {
			count++
		}

		for count >= k {
			result += int64(size - endIdx)
			if nums[startIdx] == maxNum {
				count--
			}
			startIdx++
		}

		endIdx++
	}

	return result
}