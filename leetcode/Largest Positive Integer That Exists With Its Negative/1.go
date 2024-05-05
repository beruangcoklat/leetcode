func findMaxK(nums []int) int {
	m := make(map[int]struct{})
	for _, n := range nums {
		m[n] = struct{}{}
	}

	max := -1
	for _, n := range nums {
		_, ok := m[-n]
		if n > max && ok {
			max = n
		}
	}

	return max
}