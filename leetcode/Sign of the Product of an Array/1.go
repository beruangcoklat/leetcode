func minimize(x int) int {
	if x > 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	return -1
}

func arraySign(nums []int) int {
	ret := 1
	for _, num := range nums {
		if num == 0 {
			return 0
		}
		ret *= minimize(num)
	}
	return ret
}
