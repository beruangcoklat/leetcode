func getDigits(x int) []int {
	ret := []int{}
	for x > 0 {
		digit := x % 10
		x /= 10
		ret = append([]int{digit}, ret...)
	}

	return ret
}

func addToArrayForm(num []int, k int) []int {
	lastIdx := len(num) - 1
	num[lastIdx] += k
	carrier := 0

	for i := lastIdx; i >= 0; i-- {
		carrier = num[i] / 10
		if i-1 >= 0 {
			num[i-1] += carrier
		}
		num[i] %= 10
	}

	if carrier > 0 {
		num = append(getDigits(carrier), num...)
	}

	return num
}
