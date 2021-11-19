func rangeBitwiseAnd(left int, right int) int {
	if left == right {
		return left
	}

	ret := left & (left + 1)
	for i := left + 2; i <= right; i++ {
		ret = ret & i
	}

	return ret
}
