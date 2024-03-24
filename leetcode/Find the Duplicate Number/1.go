func findDuplicate(nums []int) int {
	flag := make([]int, len(nums))
	for _, n := range nums {
		if flag[n] == 1 {
			return n
		}
		flag[n] = 1
	}
	return -1
}
