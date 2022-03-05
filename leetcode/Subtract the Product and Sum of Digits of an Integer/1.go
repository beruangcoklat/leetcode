func subtractProductAndSum(n int) int {
	if n < 10 {
		return 0
	}

	mod := 1
	sum := 0
	product := 1

	for mod <= n {
		digit := n % (mod * 10) / mod
		mod *= 10
		sum += digit
		product *= digit
	}

	return product - sum
}
