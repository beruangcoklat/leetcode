func reverse(x int) int {
	isMin := x < 0
	if isMin {
		x *= -1
	}

	result := int32(0)
	divisor := 1
	for divisor <= x {
		digit := (x / divisor) % 10
		if result > math.MaxInt32/10 || (result == math.MaxInt32/10 && digit > 7) {
			return 0
		}
		if isMin && -result < math.MinInt32/10 || (-result == math.MinInt32/10 && digit >= 8) {
			return 0
		}
		result = (result * 10) + int32(digit)
		divisor *= 10
	}

	if isMin {
		result *= -1
	}
	return int(result)
}
