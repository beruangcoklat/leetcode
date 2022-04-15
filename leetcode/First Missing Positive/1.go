func firstMissingPositive(nums []int) int {
	visited := make(map[int]struct{})
	for _, n := range nums {
		if n <= 0 {
			continue
		}
		visited[n] = struct{}{}
	}

	i := 1
	for {
		if _, ok := visited[i]; !ok {
			return i
		}
		i++
	}
}
