func isPowerOfThree(n int) bool {
	m := make(map[int]struct{})
	var a int64 = 1

	for true {
		m[int(a)] = struct{}{}
		a *= 3
		if a > math.MaxInt32-1 || int(a) >= n {
			break
		}
	}

	_, ok := m[n]
	return ok
}
