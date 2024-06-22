func minMoves(target int, maxDoubles int) int {
	steps := 0
	for target != 1 {
		if maxDoubles == 0 {
			steps += target - 1
			break
		}
		if target%2 == 0 && maxDoubles > 0 {
			target /= 2
			maxDoubles--
		} else {
			target--
		}
		steps++
	}
	return steps
}
