func calPoints(ops []string) int {
	stack := []int{}
	for _, o := range ops {
		if o == "+" {
			lastIdx := len(stack) - 1
			stack = append(stack, stack[lastIdx]+stack[lastIdx-1])
			continue
		}

		if o == "D" {
			lastIdx := len(stack) - 1
			stack = append(stack, stack[lastIdx]*2)
			continue
		}

		if o == "C" {
			lastIdx := len(stack) - 1
			stack = stack[:lastIdx]
			continue
		}

		n, _ := strconv.Atoi(o)
		stack = append(stack, n)
	}

	ret := 0
	for _, n := range stack {
		ret += n
	}
	return ret
}
