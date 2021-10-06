func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1

	for l < r {
		total := numbers[l] + numbers[r]
		if total == target {
			return []int{l + 1, r + 1}
		}

		if total > target {
			r--
		} else {
			l++
		}
	}

	return nil
}
