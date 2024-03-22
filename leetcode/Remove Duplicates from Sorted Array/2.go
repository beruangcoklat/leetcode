func removeDuplicates(nums []int) int {
	numberMap := map[int]struct{}{}
	idx := 0
	for _, num := range nums {
		_, ok := numberMap[num]
		if ok {
			continue
		}
		nums[idx] = num
		numberMap[num] = struct{}{}
		idx++
	}
	return idx
}