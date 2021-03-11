func maxSatisfied(customers []int, grumpy []int, X int) int {
	sum := 0
	for i := 0; i < len(customers); i++ {
		if grumpy[i] == 1 {
			continue
		}
		sum += customers[i]
	}

	localMax := 0
	for i := 0; i < X; i++ {
		if grumpy[i] == 1 {
			localMax += customers[i]
		}
	}

	globalMax := localMax

	for i := X; i < len(customers); i++ {
		curr := localMax
		if grumpy[i] == 1 {
			curr += customers[i]
		}
		if grumpy[i-X] == 1 {
			curr -= customers[i-X]
		}
		localMax = curr
		if localMax > globalMax {
			globalMax = localMax
		}
	}

	return globalMax + sum
}
