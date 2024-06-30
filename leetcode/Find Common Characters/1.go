func commonChars(words []string) []string {
	wordMaps := make([]map[byte]int, len(words))
	for i, word := range words {
		wordMap := make(map[byte]int)
		for _, c := range word {
			wordMap[byte(c)]++
		}
		wordMaps[i] = wordMap
	}

	result := make([]string, 0)
	for char, count := range wordMaps[0] {
		min := count
		for i := 1; i < len(wordMaps); i++ {
			currCount := wordMaps[i][char]
			if currCount < min {
				min = currCount
			}
		}
		for i := 0; i < min; i++ {
			result = append(result, string(char))
		}
	}
	return result
}