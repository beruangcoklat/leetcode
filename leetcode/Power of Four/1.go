func isPowerOfFour(n int) bool {
	var a int64 = 1
	m := map[int]struct{}{}
	for true {
		m[int(a)] = struct{}{}
		a *= 4
		if a > math.MaxInt32-1 {
			break
		}
	}

	_, ok := m[n]
	return ok
}
