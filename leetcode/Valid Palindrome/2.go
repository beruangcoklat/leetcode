func valid(c byte) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func isSame(a, b byte) bool {
	if a >= 'A' && a <= 'Z' {
		a += 'a' - 'A'
	}
	if b >= 'A' && b <= 'Z' {
		b += 'a' - 'A'
	}
	return a == b
}

func isPalindrome(s string) bool {
	size := len(s)
	left := 0
	right := size - 1

	for left < right && left < size && right >= 0 {
		if !valid(s[left]) {
			left++
			continue
		}
		if !valid(s[right]) {
			right--
			continue
		}

		if !isSame(s[left], s[right]) {
			return false
		}

		left++
		right--
	}
	return true
}