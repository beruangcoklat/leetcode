func minSubArrayLen(target int, nums []int) int {
	size := len(nums)
	leftIdx := 0
	rightIdx := 0

	minLength := 0
	minFound := false

	sum := nums[leftIdx]

	for {
		if sum >= target {
			length := rightIdx - leftIdx + 1
			if length < minLength || !minFound {
				minFound = true
				minLength = length
			}
			sum -= nums[leftIdx]
			leftIdx++
			continue
		}

		rightIdx++
		if rightIdx < size {
			sum += nums[rightIdx]
		} else {
			break
		}
	}

	return minLength
}
