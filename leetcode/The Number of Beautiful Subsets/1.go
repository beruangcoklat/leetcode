var result int

func beautifulSubsets(nums []int, k int) int {
	result = 0
	solve(nums, []int{}, -1, k)
	return result
}

func solve(nums []int, localResult []int, idx int, k int) {
	for i := idx + 1; i < len(nums); i++ {
		arr := clone(localResult, nums[i])
		if isBeautifulSubsets(arr, k) {
			result++
		}
		solve(nums, arr, i, k)
	}
}

func clone(arr []int, a int) []int {
	result := make([]int, len(arr)+1)
	copy(result, arr)
	result[len(arr)] = a
	return result
}

func isBeautifulSubsets(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if abs(nums[i]-nums[j]) == k {
				return false
			}
		}
	}
	return true
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
