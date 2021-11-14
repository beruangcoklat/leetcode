func recur(idx int, candidates []int, target int, path string, total int) [][]int {
	if total == target {
		splits := strings.Split(path, "-")
		ret := make([]int, len(splits))
		for i, s := range splits {
			ret[i], _ = strconv.Atoi(s)
		}
		return [][]int{ret}
	}

	ret := [][]int{}
	prev := -1
	first := true

	for i := idx + 1; i < len(candidates); i++ {
		if total+candidates[i] > target {
			continue
		}
		if !first && prev == candidates[i] {
			continue
		}

		first = false
		prev = candidates[i]

		var newPath string
		if path == "" {
			newPath = strconv.Itoa(candidates[i])
		} else {
			newPath = path + "-" + strconv.Itoa(candidates[i])
		}
		curr := recur(i, candidates, target, newPath, total+candidates[i])
		ret = append(ret, curr...)
	}

	return ret
}

func combinationSum2(candidates []int, target int) [][]int {
	sort.SliceStable(candidates, func(i, j int) bool {
		return candidates[i] < candidates[j]
	})
	return recur(-1, candidates, target, "", 0)
}
