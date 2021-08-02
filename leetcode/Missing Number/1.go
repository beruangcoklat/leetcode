func missingNumber(nums []int) int {
	mark := map[int]struct{}{}
	idx := 0

	for _, n := range nums {
		mark[n] = struct{}{}

		for {
			_, ok := mark[idx]
			if !ok {
				break
			}
			idx++
		}
	}

	return idx
}
