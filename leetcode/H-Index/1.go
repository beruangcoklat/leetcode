func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func hIndex(citations []int) int {
	sort.SliceStable(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	maxH := 0
	for i := 0; i < len(citations); i++ {
		paperWithMinimumCitation := len(citations) - i
		min := getMin(paperWithMinimumCitation, citations[i])
		if min > maxH {
			maxH = min
		}
	}

	return maxH
}