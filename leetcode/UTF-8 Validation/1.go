func getBinary(n int) string {
	ret := ""
	for n > 0 {
		x := n % 2
		n = n >> 1
		ret = fmt.Sprintf("%d%s", x, ret)
	}
	return fmt.Sprintf("%0*s", 8, ret)
}

func countOne(s string) int {
	ret := 0
	for _, c := range s {
		if c == '1' {
			ret++
		} else {
			break
		}
	}
	return ret
}

func validUtf8(data []int) bool {
	count := 0
	for _, d := range data {
		bin := getBinary(d)
		co := countOne(bin)
		if count > 0 {
			if co != 1 {
				return false
			}
			count--
			continue
		}

		if co == 0 {
			continue
		}

		if co == 1 || co > 4 {
			return false
		}

		count = co - 1
	}

	return count == 0
}
