func maxNumEdgesToRemove(n int, edges [][]int) int {
	aliceGraph := make([]int, n+1)
	bobGraph := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		aliceGraph[i] = -1
		bobGraph[i] = -1
	}

	minEdgeNeeded := 0
	for edgeType := 3; edgeType >= 1; edgeType-- {
		for _, edge := range edges {
			et := edge[0]
			if et != edgeType {
				continue
			}

			src := edge[1]
			dst := edge[2]
			switch edgeType {
			case 1:
				minEdgeNeeded += connect(src, dst, aliceGraph)
			case 2:
				minEdgeNeeded += connect(src, dst, bobGraph)
			case 3:
				temp1 := connect(src, dst, aliceGraph)
				temp2 := connect(src, dst, bobGraph)
				if temp1 == 1 || temp2 == 1 {
					minEdgeNeeded++
				}
			}
		}
	}

	alice := 0
	bob := 0
	for i := 1; i < n+1; i++ {
		if aliceGraph[i] == -1 {
			alice++
		}
		if bobGraph[i] == -1 {
			bob++
		}
		if alice > 1 || bob > 1 {
			return -1
		}
	}
	return len(edges) - minEdgeNeeded
}

func connect(a, b int, graph []int) int {
	aParent := findParent(a, graph)
	bParent := findParent(b, graph)
	if aParent == bParent {
		return 0
	}
	graph[aParent] = bParent
	return 1
}

func findParent(node int, graph []int) int {
	curr := node
	for graph[curr] != -1 {
		if graph[graph[curr]] != -1 {
			graph[curr] = graph[graph[curr]]
		}
		curr = graph[curr]
	}
	return curr
}
