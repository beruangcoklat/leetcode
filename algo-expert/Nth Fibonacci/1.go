package program

func GetNthFib(n int) int {
	if n <= 1 {
		return 0
	}

	if n == 2 {
		return 1
	}

	a := 0
	b := 1
	for i := 3; i <= n; i++ {
		b += a
		a = b - a
	}
	return b
}
