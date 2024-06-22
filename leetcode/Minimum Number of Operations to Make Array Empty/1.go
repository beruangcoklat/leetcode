var cache [1000001]int

func minOperations(nums []int) int {
	cache = [1000001]int{}
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	ret := 0
	for _, count := range numMap {
		op := solve(count)
		if op == -1 {
			return -1
		}
		ret += op
	}
	return ret
}

func solve(n int) int {
	if n == 0 {
		return 0
	}

	if cache[n] != 0 {
		return cache[n]
	}

	min := -1
	if n-3 >= 0 && n-3 != 1 {
		curr := solve(n-3) + 1
		if min == -1 || curr < min {
			min = curr
		}
	}
	if n-2 >= 0 && n-2 != 1 {
		curr := solve(n-2) + 1
		if min == -1 || curr < min {
			min = curr
		}
	}

	cache[n] = min
	return min
}