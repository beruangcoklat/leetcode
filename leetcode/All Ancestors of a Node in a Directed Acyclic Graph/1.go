type node struct {
	id      int
	parents []*node
}

func getAncestors(n int, edges [][]int) [][]int {
	graph := make(map[int]*node)
	for i := 0; i < n; i++ {
		graph[i] = &node{
			id:      i,
			parents: make([]*node, 0),
		}
	}

	for _, edge := range edges {
		src := edge[0]
		dst := edge[1]
		graph[dst].parents = append(graph[dst].parents, graph[src])
	}

	result := make([][]int, n)
	for i := 0; i < n; i++ {
		if len(graph[i].parents) == 0 {
			continue
		}
		m := make(map[int]struct{})
		solve(graph[i], m)
		result[i] = mapToSortedArray(m)
	}
	return result
}

func solve(n *node, m map[int]struct{}) {
	for i := 0; i < len(n.parents); i++ {
		id := n.parents[i].id
		_, ok := m[id]
		if ok {
			continue
		}
		m[id] = struct{}{}
		solve(n.parents[i], m)
	}
}

func mapToSortedArray(m map[int]struct{}) []int {
	arr := make([]int, len(m))
	i := 0
	for n := range m {
		arr[i] = n
		i++
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return arr
}
