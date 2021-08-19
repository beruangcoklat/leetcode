var memo map[string]int

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solve(prices []int, idx, tx, k int, isBuy bool) int {
	if tx == k || idx == len(prices) {
		return 0
	}

	key := fmt.Sprintf("%v-%v-%v", idx, tx, isBuy)
	if mem, ok := memo[key]; ok {
		return mem
	}

	keep := solve(prices, idx+1, tx, k, isBuy)
	curr := 0
	if isBuy {
		curr = solve(prices, idx+1, tx, k, !isBuy) - prices[idx]
	} else {
		curr = solve(prices, idx+1, tx+1, k, !isBuy) + prices[idx]
	}

	max := getMax(keep, curr)
	memo[key] = max
	return max
}

func maxProfit(k int, prices []int) int {
	memo = make(map[string]int)
	return solve(prices, 0, 0, k, true)
}
