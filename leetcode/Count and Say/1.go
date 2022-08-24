func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}

	ret := ""
	prev := ""
	count := 0
	for _, c := range countAndSay(n - 1) {
		s := string(c)
		if s != prev {
			if count > 0 {
				ret += fmt.Sprintf("%v%v", count, prev)
			}
			prev = s
			count = 1
		} else {
			count++
		}
	}
	ret += fmt.Sprintf("%v%v", count, prev)

	return ret
}
