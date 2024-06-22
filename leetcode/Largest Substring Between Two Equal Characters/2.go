func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[rune][]int)
	for i, c := range s {
		m[c] = append(m[c], i)
	}

	max := -1
	for _, indexs := range m {
		if len(indexs) < 2 {
			continue
		}

		curr := indexs[len(indexs)-1] - indexs[0] - 1
		if curr > max {
			max = curr
		}
	}

	return max
}
