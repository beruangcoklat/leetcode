func max(n ...int) int {
	max := n[0]
	for i := 1; i < len(n); i++ {
		if max < n[i] {
			max = n[i]
		}
	}
	return max
}

var memo map[string][]int

func solve(isAlex bool, piles []int, right, left int) (int, int) {
	if len(piles) == 1 {
		if isAlex {
			return piles[0], 0
		}
		return 0, piles[0]
	}

	key := fmt.Sprintf("%v-%v-%v", isAlex, right, left)
	mem, ok := memo[key]
	if ok {
		return mem[0], mem[1]
	}

	alex, lee := 0, 0

	a, b := solve(!isAlex, piles[1:], right+1, left)
	c, d := solve(!isAlex, piles[:len(piles)-1], right, left+1)

	if isAlex {
		a += piles[0]
		c += piles[len(piles)-1]
		alex = max(a, c)
	} else {
		b += piles[0]
		d += piles[len(piles)-1]
		lee = max(b, d)
	}

	memo[key] = []int{alex, lee}
	return alex, lee
}

func stoneGame(piles []int) bool {
	memo = make(map[string][]int)
	a, b := solve(true, piles, 0, 0)
	return a > b
}
