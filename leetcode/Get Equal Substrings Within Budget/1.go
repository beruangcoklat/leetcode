func equalSubstring(s string, t string, maxCost int) int {
	l, r := -1, -1
	diff := make([]int, len(s))

	for i := 0; i < len(s); i++ {
		diff[i] = abs(int(s[i]) - int(t[i]))
		if l == -1 && diff[i] <= maxCost {
			l = i
			r = i
		}
	}

	if l == -1 {
		return 0
	}

	currDiff := diff[l]
	maxLength := 1
	for r < len(s) {
		if currDiff <= maxCost {
			currLength := r - l + 1
			if currLength > maxLength {
				maxLength = currLength
			}
		}

		if currDiff > maxCost {
			currDiff -= diff[l]
			l++
		} else {
			r++
			if r < len(s) {
				currDiff += diff[r]
			}
		}
	}
	return maxLength
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
