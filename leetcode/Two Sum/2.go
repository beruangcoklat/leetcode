func twoSum(nums []int, target int) []int {
	mark := map[int]int{}
	for i, n := range nums {
		idx, ok := mark[target-n]
		if ok {
			return []int{idx, i}
		}
		mark[n] = i
	}
	return nil
}
