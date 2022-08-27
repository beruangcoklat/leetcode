func runeToInt(a byte) int {
	return int(a) - '0'
}

func addStrings(a string, b string) string {
	i := 1
	res := []int{}
	carry := 0

	for len(a)-i >= 0 && len(b)-i >= 0 {
		aa := a[len(a)-i]
		bb := b[len(b)-i]

		x := runeToInt(aa) + runeToInt(bb) + carry
		carry = 0
		if x > 9 {
			x -= 10
			carry += 1
		}
		res = append([]int{x}, res...)
		i++
	}

	for len(a)-i >= 0 {
		x := runeToInt(a[len(a)-i]) + carry
		carry = 0
		if x > 9 {
			x -= 10
			carry += 1
		}
		res = append([]int{x}, res...)
		i++
	}

	for len(b)-i >= 0 {
		x := runeToInt(b[len(b)-i]) + carry
		carry = 0
		if x > 9 {
			x -= 10
			carry += 1
		}
		res = append([]int{x}, res...)
		i++
	}

	if carry > 0 {
		res = append([]int{carry}, res...)
	}

	strRes := ""
	for _, a := range res {
		strRes += strconv.Itoa(a)
	}
	return strRes
}
