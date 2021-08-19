func maxProfit(prices []int) int {
	maxProfit := 0
	idx := 0
	for idx+1 < len(prices) {
		for idx+1 < len(prices) && prices[idx+1] < prices[idx] {
			idx++
		}
		lowest := prices[idx]

		for idx+1 < len(prices) && prices[idx+1] > prices[idx] {
			idx++
		}
		highest := prices[idx]
		maxProfit += highest - lowest
		idx++
	}
	return maxProfit
}
