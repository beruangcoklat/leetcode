func distributeCandies(candyType []int) int {
	mark := map[int]struct{}{}
	uniqueCount := 0

	for _, ct := range candyType {
		_, ok := mark[ct]
		if ok {
			continue
		}
		mark[ct] = struct{}{}
		uniqueCount++
	}

	max := len(candyType) / 2

	if uniqueCount > max {
		return max
	}

	return uniqueCount
}
