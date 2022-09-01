func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func trap(height []int) int {
	total := 0
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))

	max := -1
	for i := 0; i < len(height); i++ {
		max = getMax(max, height[i])
		leftMax[i] = max
	}

	max = -1
	for i := len(height) - 1; i >= 0; i-- {
		max = getMax(max, height[i])
		rightMax[i] = max
	}

	for i := 1; i < len(height)-1; i++ {
		curr := getMin(leftMax[i-1], rightMax[i+1]) - height[i]
		curr = getMax(curr, 0)
		total += curr
	}
	return total
}
