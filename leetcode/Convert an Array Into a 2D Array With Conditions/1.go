func findMatrix(nums []int) [][]int {
	result := make([][]int, 0)
	for len(nums) > 0 {
		numMap := make(map[int]struct{})
		newNums := []int{}
		row := []int{}
		for _, num := range nums {
			_, ok := numMap[num]
			if ok {
				newNums = append(newNums, num)
				continue
			}
			numMap[num] = struct{}{}
			row = append(row, num)
		}
		result = append(result, row)
		nums = newNums
	}
	return result
}