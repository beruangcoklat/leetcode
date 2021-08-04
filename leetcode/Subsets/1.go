var (
	ret [][]int
)

func cloneArr(arr []int) []int {
	result := make([]int, len(arr))
	for i, v := range arr {
		result[i] = v
	}
	return result
}

func recur(idx, n int, arr, result []int) {
    ret = append(ret, result)
	for i := idx + 1; i < len(arr); i++ {
		clone := cloneArr(result)
		clone = append(clone, arr[i])
		recur(i, n, arr, clone)
	}
}

func subsets(nums []int) [][]int {
	ret = [][]int{}
	recur(-1, len(nums), nums, []int{})
	return ret
}
