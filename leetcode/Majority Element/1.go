func majorityElement(nums []int) int {
	countMap := make(map[int]int)
	for _, num := range nums {
		countMap[num]++
	}

	threshold := len(nums) / 2
	ret := 0
	maxCount := 0

	for num, count := range countMap {
		if count <= threshold {
			continue
		}

		if count > maxCount {
			maxCount = count
			ret = num
		}
	}
	return ret
}