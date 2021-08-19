func maxProfit(prices []int) int {
	max := 0
	maxProfit := 0
	for i := len(prices) - 1; i >= 0; i-- {
		if i == len(prices)-1 || prices[i] > max {
			max = prices[i]
		}
		currMaxProfit := max - prices[i]
		if currMaxProfit > maxProfit {
			maxProfit = currMaxProfit
		}
	}
	return maxProfit
}
