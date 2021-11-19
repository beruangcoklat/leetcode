func isHappy(n int) bool {
	for true {
		if n == 1 || n == 7 {
			return true
		}

		if n < 10 {
			return false
		}

		mul := 1
		total := 0

		for true {
			if mul > n {
				break
			}
			digit := n / mul % 10
			mul *= 10
			total += (digit * digit)
		}

		n = total
	}
	return false
}
