func singleNumber(nums []int) int {
	numMap := make(map[int]int)
	for _, n := range nums {
		numMap[n]++
	}
	for num, count := range numMap {
		if count == 1 {
			return num
		}
	}
	return 0
}
