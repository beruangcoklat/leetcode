func isHappy(n int) bool {
	m := make(map[int]int)
	for {
		if n == 1 {
			return true
		}

		if n < 10 {
			m[n]++
			if m[n] > 1 {
				return false
			}
		}

		total := 0
		for n > 0 {
			num := n % 10
			total += (num * num)
			n /= 10
		}
		n = total
	}
}
