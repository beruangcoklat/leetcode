func removeDuplicates(nums []int) int {
	numberMap := map[int]int{}
	idx := 0
	for _, num := range nums {
		count := numberMap[num]
		if count >= 2 {
			continue
		}
		nums[idx] = num
		numberMap[num]++
		idx++
	}
	return idx
}