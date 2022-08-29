func merge(intervals [][]int) [][]int {
	for i := 0; i < len(intervals); i++ {
		merged := false
		for j := i + 1; j < len(intervals); j++ {
			if !isOverlap(intervals[i], intervals[j]) {
				continue
			}

			merged = true
			intervals[i] = mergeIntervals(intervals[i], intervals[j])
			intervals = append(intervals[:j], intervals[j+1:]...)
			break
		}

		if merged {
			i--
		}
	}
	return intervals
}

func getLeftAndRight(a []int, b []int) ([]int, []int) {
	var left, right []int

	if a[0] == b[0] {
		if a[1] < b[1] {
			left = a
			right = b
		} else {
			left = b
			right = a
		}
	} else if a[0] < b[0] {
		left = a
		right = b
	} else {
		left = b
		right = a
	}

	return left, right
}

func isOverlap(a []int, b []int) bool {
	left, right := getLeftAndRight(a, b)
	return left[1] >= right[0]
}

func mergeIntervals(a []int, b []int) []int {
	min := getMin(a[0], b[0])
	max := getMax(a[1], b[1])
	return []int{min, max}
}

func getMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func getMax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
