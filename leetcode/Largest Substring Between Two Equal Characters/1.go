func maxLengthBetweenEqualCharacters(s string) int {
	max := -1
	for start := 0; start < len(s); start++ {
		for end := start + 1; end < len(s); end++ {
			if s[start] != s[end] {
				continue
			}

			curr := end - start - 1
			if curr > max {
				max = curr
			}
		}
	}
	return max
}