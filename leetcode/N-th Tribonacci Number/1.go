func tribonacci(n int) int {
	if n <= 1 {
		return n
	}
	if n == 2 {
		return 1
	}
	a, b, c := 0, 1, 1
	curr := 0
	for i := 3; i <= n; i++ {
		curr = a + b + c
		a = b
		b = c
		c = curr
	}
	return curr
}
