func findErrorNums(nums []int) []int {
	doubleNum := 0
	missingNum := 0

	mark := map[int]struct{}{}
	for _, n := range nums {
		_, ok := mark[n]
		if ok {
			doubleNum = n
			continue
		}
		mark[n] = struct{}{}
	}

	for i := 1; i <= len(nums); i++ {
		_, ok := mark[i]
		if !ok {
			missingNum = i
			break
		}
	}

	return []int{doubleNum, missingNum}
}
