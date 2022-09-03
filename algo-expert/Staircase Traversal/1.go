package main

var cache map[int]int

func solve(height int, maxSteps int) int {
	if height == 0 {
		return 1
	}

	val, ok := cache[height]
	if ok {
		return val
	}

	total := 0
	for i := 1; i <= maxSteps; i++ {
		if height-i >= 0 {
			total += solve(height-i, maxSteps)
		}
	}

	cache[height] = total
	return total
}

func StaircaseTraversal(height int, maxSteps int) int {
	cache = make(map[int]int)
	return solve(height, maxSteps)
}
