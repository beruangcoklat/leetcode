func canCompleteCircuit(gas []int, cost []int) int {
	size := len(gas)
	diff := 0

	total := 0
	ret := 0
	for i := 0; i < size; i++ {
		diff += gas[i] - cost[i]
		total += gas[i] - cost[i]
		if total < 0 {
			total = 0
			ret = i + 1
		}
	}

	if diff < 0 {
		return -1
	}

	return ret
}
