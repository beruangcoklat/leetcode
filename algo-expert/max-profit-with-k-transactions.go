package program

var memo [][][]int

func solve(prices []int, l int, k int, hold bool, holdValueIdx int, curr int) int {
	hidx := 0
	if hold {
		hidx = 1
	}

	if curr >= l {
		return 0
	}

	if k == 0 {
		return 0
	}

	if memo[k-1][hidx][curr] != 0 {
		return memo[k-1][hidx][curr]
	}

	if curr == l-1 {
		if hold && prices[curr] > prices[holdValueIdx] {
			return prices[curr] - prices[holdValueIdx]
		}
		return 0
	}

	max := -99

	if hold {
		for i := curr; i < l; i++ {
			if prices[i] > prices[holdValueIdx] {
				s := solve(prices, l, k-1, false, -1, i+1) + (prices[i] - prices[holdValueIdx])
				if s > max {
					max = s
				}
			}
		}
	} else {
		ifhold := solve(prices, l, k, true, curr, curr+1)
		ifnothold := solve(prices, l, k, false, -1, curr+1)
		if ifhold > ifnothold {
			max = ifhold
		} else {
			max = ifnothold
		}
	}

	memo[k-1][hidx][curr] = max
	return max
}

func MaxProfitWithKTransactions(prices []int, k int) int {
	l := len(prices)
	memo = make([][][]int, k)
	for a := 0; a < k; a++ {
		memo[a] = make([][]int, 2)
		for b := 0; b < 2; b++ {
			memo[a][b] = make([]int, l)
		}
	}
	return solve(prices, l, k, false, -1, 0)
}
