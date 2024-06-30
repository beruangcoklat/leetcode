func isNStraightHand(hand []int, groupSize int) bool {
	m := make(map[int]int)
	for _, n := range hand {
		m[n]++
	}
	sort.SliceStable(hand, func(i, j int) bool {
		return hand[i] < hand[j]
	})

	count := 0
	for i := 0; i < len(hand); i++ {
		valid := true
		for j := hand[i]; j < hand[i]+groupSize; j++ {
			if m[j] == 0 {
				valid = false
				break
			}
		}
		if !valid {
			continue
		}
		for j := hand[i]; j < hand[i]+groupSize; j++ {
			m[j]--
			if m[j] == 0 {
				delete(m, j)
			}
		}
		count++
	}

	return len(m) == 0
}