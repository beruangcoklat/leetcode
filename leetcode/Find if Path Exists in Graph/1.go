func validPath(n int, edges [][]int, source int, destination int) bool {
	edgeMap := make(map[int][]int)
	for _, edge := range edges {
		src := edge[0]
		dst := edge[1]
		edgeMap[src] = append(edgeMap[src], dst)
		edgeMap[dst] = append(edgeMap[dst], src)
	}

	list := []int{source}
	visitedMap := make(map[int]struct{})
	for len(list) > 0 {
		curr := list[0]
		if curr == destination {
			return true
		}

		list = list[1:]
		_, ok := visitedMap[curr]
		if ok {
			continue
		}
		visitedMap[curr] = struct{}{}

		list = append(list, edgeMap[curr]...)
	}

	return false
}
