var (
	result = [][]int{}
)

func cloneArr(arr []int) []int {
	ret := make([]int, len(arr))
	for i, a := range arr {
		ret[i] = a
	}
	return ret
}

func cloneVisited(visited map[int]struct{}) map[int]struct{} {
	ret := make(map[int]struct{})
	for k := range visited {
		ret[k] = struct{}{}
	}
	return ret
}

func recur(idx, k, n int, arr []int, visited map[int]struct{}) {
	if len(arr) == k {
		result = append(result, arr)
		return
	}

	for i := idx + 1; i <= n; i++ {
		if _, ok := visited[i]; ok {
			continue
		}
		newVisited := cloneVisited(visited)
		newVisited[i] = struct{}{}
		newArr := cloneArr(arr)
		newArr = append(newArr, i)
		recur(i, k, n, newArr, newVisited)
	}
}

func combine(n int, k int) [][]int {
	result = [][]int{}
	recur(0, k, n, []int{}, make(map[int]struct{}))
	return result
}
