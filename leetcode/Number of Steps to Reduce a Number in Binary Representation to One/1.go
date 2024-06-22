func numSteps(s string) int {
	arr := make([]bool, len(s)+1)
	for i, c := range s {
		arr[i] = c == '1'
	}

	steps := 0
	idx := len(s) - 1
	for idx > 0 {
		if arr[idx] {
			arr[idx] = false
			breaked := false
			for i := idx - 1; i > 0; i-- {
				if !arr[i] {
					arr[i] = true
					breaked = true
					break
				} else {
					arr[i] = false
				}
			}
			if !breaked {
				idx++
				arr[idx] = false
			}
			idx++
		}
		steps++
		idx--
	}
	return steps
}