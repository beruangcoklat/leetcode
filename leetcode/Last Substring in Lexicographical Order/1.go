func lastSubstring(s string) string {
	var max rune
	first := true
	for _, a := range s {
		if first || a > max {
			max = a
		}
		first = false
	}
	maxStr := string(max)

	ret := ""
	for i := 0; i < len(s); i++ {
		if s[i:i+1] != maxStr {
			continue
		}

		length := len(s) - i
		curr := s[i : i+length]
		if ret == "" || curr > ret {
			ret = s[i : i+length]
		}
	}
	return ret
}
