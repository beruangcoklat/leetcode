func lengthOfLIS(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = 1
	for i := 1; i < len(nums); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] && dp[j]+1 > dp[i] {
				dp[i] = dp[j] + 1
			}
		}
	}

	max := dp[0]
	for i := 1; i < len(dp); i++ {
		if dp[i] > max {
			max = dp[i]
		}
	}
	return max
}
