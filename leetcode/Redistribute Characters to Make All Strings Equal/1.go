func makeEqual(words []string) bool {
	m := make(map[rune]int)
	for _, word := range words {
		for _, c := range word {
			m[c]++
		}
	}
	for _, count := range m {
		if count%len(words) != 0 {
			return false
		}
	}
	return true
}
