func countOdds(low int, high int) int {
	r := high - low + 1
	result := r / 2
	if r%2 == 1 && low%2 == 1 {
		result++
	}
	return result
}
