func getMin(n ...int) int {
	min := n[0]
	for i := 1; i < len(n); i++ {
		if min > n[i] {
			min = n[i]
		}
	}
	return min
}

var memo map[string]int

func solve(word1 string, word2 string) int {

	if len(word1) == 0 {
		return len(word2)
	}

	if len(word2) == 0 {
		return len(word1)
	}

	key := word1 + "###" + word2
	val, ok := memo[key]
	if ok {
		return val
	}

	var max int

	if word1[:1] == word2[:1] {
		max = solve(word1[1:], word2[1:])
	} else {
		insert := solve(word1, word2[1:]) + 1
		update := solve(word1[1:], word2[1:]) + 1
		delete := solve(word1[1:], word2) + 1
		max = getMin(insert, update, delete)
	}

	memo[key] = max
	return max
}

func minDistance(word1 string, word2 string) int {
	memo = make(map[string]int)
	return solve(word1, word2)
}
