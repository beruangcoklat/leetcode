// https://leetcode.com/problems/reverse-integer/

func reverse(x int) int {
	min := -2147483648
	max := 2147483647

	minus := false
	if x < 0 {
		minus = true
		x *= -1
	}
	ret := 0
	multi := 10
	stop := 0
	for stop < 2 {
		if x/multi == 0 {
			stop++
		}
		if stop == 2 {
			break
		}
		ret *= 10
		ret += (x % multi / (multi / 10))

		if ret > max {
			return 0
		}
		if minus && (ret*-1) < min {
			return 0
		}

		multi *= 10
	}
	if minus {
		ret *= -1
	}
	return ret
}