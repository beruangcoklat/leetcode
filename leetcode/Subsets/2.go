var result [][]int

func subsets(nums []int) [][]int {
	result = make([][]int, 0)
	result = append(result, []int{})
	solve(nums, []int{}, -1)
	return result
}

func solve(nums []int, localResult []int, idx int) {
	for i := idx + 1; i < len(nums); i++ {
		arr := clone(localResult, nums[i])
		result = append(result, arr)
		solve(nums, arr, i)
	}
}

func clone(arr []int, a int) []int {
	result := make([]int, len(arr)+1)
	copy(result, arr)
	result[len(arr)] = a
	return result
}