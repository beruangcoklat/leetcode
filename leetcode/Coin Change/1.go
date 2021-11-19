type data struct {
	val   int
	valid bool
}

var memo map[string]data

func solve(idx int, coins []int, amount int) data {
	if amount == 0 {
		return data{valid: true}
	}

	key := fmt.Sprintf("%v-%v", idx, amount)
	if val, ok := memo[key]; ok {
		return val
	}

	min := data{}
	for i := idx; i < len(coins); i++ {
		if amount-coins[i] < 0 {
			continue
		}

		curr := solve(i, coins, amount-coins[i])
        if !curr.valid {
            continue
        }
		curr.val += 1
		if !min.valid || min.val > curr.val {
			min = curr
		}
	}

	memo[key] = min
	return min
}

func coinChange(coins []int, amount int) int {
	memo = make(map[string]data)
	ret := solve(0, coins, amount)
	if ret.valid {
		return ret.val
	}
	return -1
}
