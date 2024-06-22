func minOperations(nums []int) int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	ret := 0
	for _, count := range numMap {
		op := getOps(count)
		if op == -1 {
			return -1
		}
		ret += op
	}
	return ret
}

func getOps(n int) int {
	if n < 2 {
		return -1
	}
	if n == 2 {
		return 1
	}
	ret := n / 3
	if n%3 != 0 {
		ret++
	}
	return ret
}
