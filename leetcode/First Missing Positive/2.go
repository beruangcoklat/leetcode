func firstMissingPositive(nums []int) int {
	visited := make([]bool, len(nums)+1)
	for _, n := range nums {
		if n <= 0 || n > len(nums) {
			continue
		}
		visited[n] = true
	}

	for i := 1; i < len(nums)+1; i++ {
		if !visited[i] {
			return i
		}
	}
	return len(nums) + 1
}
