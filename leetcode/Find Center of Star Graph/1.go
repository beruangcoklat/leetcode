func findCenter(edges [][]int) int {
	visited := map[string]struct{}{}
	count := map[int]int{}

	for _, edge := range edges {
		src := edge[0]
		dst := edge[1]
		if src < dst {
			src, dst = dst, src
		}
		key := fmt.Sprintf("%d-%d", src, dst)
		if _, ok := visited[key]; ok {
			continue
		}

		visited[key] = struct{}{}
		count[src]++
		count[dst]++
	}

	maxC := 0
	maxNode := 0
	for node, c := range count {
		if c > maxC {
			maxC = c
			maxNode = node
		}
	}
	return maxNode
}