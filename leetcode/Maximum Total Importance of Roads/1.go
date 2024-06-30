type node struct {
	number int
	count  int
	weight int
}

func maximumImportance(n int, roads [][]int) int64 {
	countMap := make(map[int]int)
	roadMap := make(map[string]struct{})

	for _, road := range roads {
		src := road[0]
		dst := road[1]
		if src > dst {
			src, dst = dst, src
		}

		key := fmt.Sprintf("%d-%d", src, dst)
		if _, ok := roadMap[key]; ok {
			continue
		}
		roadMap[key] = struct{}{}

		countMap[src]++
		countMap[dst]++
	}

	nodes := make([]*node, 0)
	nodeMap := make(map[int]*node)

	for number, count := range countMap {
		newNode := &node{
			number: number,
			count:  count,
		}
		nodes = append(nodes, newNode)
		nodeMap[number] = newNode
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].count > nodes[j].count
	})

	for i := 0; i < len(nodes); i++ {
		nodes[i].weight = n - i
	}

	roadMap = make(map[string]struct{})
	var result int64
	for _, road := range roads {
		src := road[0]
		dst := road[1]
		if src > dst {
			src, dst = dst, src
		}

		key := fmt.Sprintf("%d-%d", src, dst)
		if _, ok := roadMap[key]; ok {
			continue
		}
		roadMap[key] = struct{}{}

		result += int64(nodeMap[src].weight + nodeMap[dst].weight)
	}
	return result
}
