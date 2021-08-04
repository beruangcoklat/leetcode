var (
	ret  [][]int
	mark map[string]struct{}
)

func toString(arr []int) string {
	str := ""
	for _, a := range arr {
		str = fmt.Sprintf("%s%d", str, a)
	}
	return str
}

func cloneArr(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}

func recur(idx, n int, arr, result []int) {
	key := toString(result)
	_, ok := mark[key]
	if !ok {
		ret = append(ret, result)
		mark[key] = struct{}{}
	}

	for i := idx + 1; i < len(arr); i++ {
		clone := cloneArr(result)
		clone = append(clone, arr[i])
		recur(i, n, arr, clone)
	}
}

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)
	ret = [][]int{}
	mark = map[string]struct{}{}
	recur(-1, len(nums), nums, []int{})
	return ret
}
