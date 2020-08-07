// https://leetcode.com/problems/permutations-ii/

var result [][]int

func newArr(arr []int, val int) []int {
	ret := []int{}
	for _, a := range arr {
		ret = append(ret, a)
	}
	ret = append(ret, val)
	return ret
}

func recursive(nums []int, idx int, visited int, arr []int) {
	if len(arr) == len(nums) {
		result = append(result, arr)
		return
	}

	dict := make(map[int]struct{})

	for i := 0; i < len(nums); i++ {
		if (visited>>i)%2 == 1 {
			continue
		}

		if _, ok := dict[nums[i]]; ok {
			continue
		}
		dict[nums[i]] = struct{}{}

		newVisited := visited | (1 << i)
		recursive(nums, i, newVisited, newArr(arr, nums[i]))
	}
}

func permuteUnique(nums []int) [][]int {
	result = [][]int{}
	recursive(nums, 0, 0, []int{})
	return result
}
