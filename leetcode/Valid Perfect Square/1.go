func isPerfectSquare(num int) bool {
	i := int64(1)
	ii := int64(1)

	for {
		ii = i * i
		if ii == int64(num) {
			return true
		}
		if ii > math.MaxInt32 {
			return false
		}
        i++
	}
}
