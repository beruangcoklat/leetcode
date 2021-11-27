func findDuplicates(nums []int) []int {
	numMap := make(map[int]int)
	ret := []int{}
	for _, n := range nums {
		numMap[n]++
		if numMap[n] > 1 {
			ret = append(ret, n)
		}
	}
	return ret
}
