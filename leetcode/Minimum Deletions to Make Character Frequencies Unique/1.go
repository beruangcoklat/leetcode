type node struct {
	char  rune
	count int
}

func minDeletions(s string) int {
	nodes := []*node{}
	countMap := make(map[rune]*node)
	for _, c := range s {
		n, ok := countMap[c]
		if !ok {
			nn := &node{
				char:  c,
				count: 1,
			}
			countMap[c] = nn
			nodes = append(nodes, nn)
			continue
		}
		n.count++
	}

	sort.SliceStable(nodes, func(i, j int) bool {
		return nodes[i].count > nodes[j].count
	})

	ret := 0
	for i := 1; i < len(nodes); i++ {
		if nodes[i].count >= nodes[i-1].count {
			var x int
			if nodes[i-1].count == 0 {
				x = nodes[i].count
			} else {
				x = (1 + nodes[i].count - nodes[i-1].count)
			}

			ret += x
			nodes[i].count -= x
		}
	}

	return ret
}
