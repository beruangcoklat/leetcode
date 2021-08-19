var memo map[string]int

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(prices []int, idx, tx int, isBuy bool) int {
	if tx == 2 || idx == len(prices) {
		return 0
	}

	key := fmt.Sprintf("%v-%v-%v", idx, tx, isBuy)
	if mem, ok := memo[key]; ok {
		return mem
	}

	keep := solve(prices, idx+1, tx, isBuy)
	curr := 0
	if isBuy {
		curr = solve(prices, idx+1, tx, !isBuy) - prices[idx]
	} else {
		curr = solve(prices, idx+1, tx+1, !isBuy) + prices[idx]
	}

	max := getMax(keep, curr)
	memo[key] = max
	return max
}

func maxProfit(prices []int) int {
	memo = make(map[string]int)
	return solve(prices, 0, 0, true)
}
