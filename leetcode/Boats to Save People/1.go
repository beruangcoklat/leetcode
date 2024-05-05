func numRescueBoats(people []int, limit int) int {
	sort.SliceStable(people, func(i, j int) bool {
		return people[i] < people[j]
	})

	left, right := 0, len(people)-1
	ret := 0
	for left <= right {
		sum := people[left] + people[right]
		if sum <= limit {
			left++
			right--
		} else {
			right--
		}
		ret++
	}
	return ret
}
