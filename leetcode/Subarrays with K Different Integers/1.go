func solve(nums []int, k int) int {
	size := len(nums)
	result := 0
	m := make(map[int]int)

	leftIdx, rightIdx := 0, 0

	for rightIdx < size {
		m[nums[rightIdx]]++

		for len(m) > k {
			m[nums[leftIdx]]--
			if m[nums[leftIdx]] == 0 {
				delete(m, nums[leftIdx])
			}
			leftIdx++
		}

		result += (rightIdx - leftIdx + 1)
		rightIdx++
	}

	return result
}

func subarraysWithKDistinct(nums []int, k int) int {
	return solve(nums, k) - solve(nums, k-1)
}
