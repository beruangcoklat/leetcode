func cloneArr(arr []int) []int {
	ret := make([]int, len(arr))
	for i, a := range arr {
		ret[i] = a
	}
	return ret
}

func recur(currNode, lastNode int, graph [][]int, path []int) [][]int {
	if currNode == lastNode {
		return [][]int{append(path, lastNode)}
	}
	ret := [][]int{}
	for _, n := range graph[currNode] {
		res := recur(n, lastNode, graph, append(cloneArr(path), currNode))
		ret = append(ret, res...)
	}
	return ret
}

func allPathsSourceTarget(graph [][]int) [][]int {
	return recur(0, len(graph)-1, graph, []int{})
}
