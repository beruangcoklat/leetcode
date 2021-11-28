func fib(n int) int {
	if n <= 1 {
		return n
	}
	a, b := 0, 1
	curr := 0
	for i := 2; i <= n; i++ {
		curr = a + b
		a = b
		b = curr
	}
	return curr
}
