// https://leetcode.com/problems/string-to-integer-atoi/

func myAtoi(str string) int {
	ret := 0
	starting := false
	ismin := false
	min := -2147483648
	max := 2147483647

	for i := 0; i < len(str); i++ {
		if !starting {
			if str[i] == '+' {
				starting = true
				continue
			} else if str[i] == '-' {
				starting = true
				ismin = true
				continue
			} else if str[i] == ' ' {
				continue
			}
		}

		if str[i] < '0' || str[i] > '9' {
			break
		}

		starting = true
		ret *= 10
		ret += int(str[i] - '0')
		if ismin && (ret*-1) <= min {
			return min
		}
		if !ismin && ret >= max {
			return max
		}
	}

	if ismin {
		ret *= -1
	}
	return ret
}