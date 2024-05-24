func findRelativeRanks(score []int) []string {
	sortedScore := make([]int, len(score))
	copy(sortedScore, score)

	sort.SliceStable(sortedScore, func(i, j int) bool {
		return sortedScore[i] > sortedScore[j]
	})

	rankMap := make(map[int]string)
	for i := 0; i < len(sortedScore); i++ {
		var value string
		if i == 0 {
			value = "Gold Medal"
		} else if i == 1 {
			value = "Silver Medal"
		} else if i == 2 {
			value = "Bronze Medal"
		} else {
			value = fmt.Sprintf("%d", i+1)
		}
		rankMap[sortedScore[i]] = value
	}

	ret := []string{}
	for _, s := range score {
		ret = append(ret, rankMap[s])
	}
	return ret
}
