func isPerfectSquare(num int) bool {
	if num > 46341*46341 {
		return false
	}

	left, right := 1, 46341

	for left <= right {
		mid := (left + right) / 2
		m := mid * mid
		if m == num {
			return true
		}

		if m > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return false
}
