func validateStackSequences(pushed []int, popped []int) bool {
	stack := []int{}
	idxPop := 0

	for _, v := range pushed {
		stack = append(stack, v)
		for len(stack) > 0 && idxPop < len(popped) && stack[len(stack)-1] == popped[idxPop] {
			stack = stack[:len(stack)-1]
			idxPop++
		}
	}

	for idxPop < len(popped) {
		if popped[idxPop] == stack[len(stack)-1] {
			idxPop++
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return true
}
