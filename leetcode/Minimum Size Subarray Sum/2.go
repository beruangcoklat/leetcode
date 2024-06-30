func minSubArrayLen(target int, nums []int) int {
	l, r := 0, 0
	sum := nums[0]
	min := 0
	if sum >= target {
		min = 1
	}

	for {
		if sum >= target {
			sum -= nums[l]
			l++
		} else {
			r++
			if r == len(nums) {
				break
			}
			sum += nums[r]
		}

		if sum < target {
			continue
		}
		curr := r - l + 1
		if min == 0 || curr < min {
			fmt.Println(nums[l : r+1])
			min = curr
		}
	}
	return min
}