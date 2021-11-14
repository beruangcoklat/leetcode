var ret [][]int

func cloneArr(arr []int) []int {
	ret := make([]int, len(arr))
	for i, a := range arr {
		ret[i] = a
	}
	return ret
}

func recur(idx int, candidates []int, target int, path []int, total int) {
	if total == target {
		ret = append(ret, path)
		return
	}

	for i := idx; i < len(candidates); i++ {
		if total+candidates[i] > target {
			continue
		}
		newPath := append(cloneArr(path), candidates[i])
		recur(i, candidates, target, newPath, total+candidates[i])
	}
}

func combinationSum(candidates []int, target int) [][]int {
	ret = [][]int{}
	recur(0, candidates, target, []int{}, 0)
	return ret
}
