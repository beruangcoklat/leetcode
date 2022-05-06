func removeDuplicates(s string, k int) string {
	stack := []rune{}
	for _, c := range s {
		stack = append(stack, c)
		if len(stack) < k {
			continue
		}

		isSame := true
		for i := len(stack) - k; i < len(stack)-1; i++ {
			if stack[i] != stack[i+1] {
				isSame = false
				break
			}
		}
		if isSame {
			stack = stack[:len(stack)-k]
		}
	}

	return string(stack)
}
