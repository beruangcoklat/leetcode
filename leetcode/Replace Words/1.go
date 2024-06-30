func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	sort.SliceStable(words, func(i, j int) bool {
		return len(words[i]) > len(words[j])
	})

	sort.SliceStable(dictionary, func(i, j int) bool {
		return len(dictionary[i]) < len(dictionary[j])
	})

	swapMap := make(map[string]string)

	for _, word := range words {
		for _, dict := range dictionary {
			if strings.HasPrefix(word, dict) {
				swapMap[word] = dict
				break
			}
		}
	}

	words = strings.Split(sentence, " ")
	for i, word := range words {
		if newWord, ok := swapMap[word]; ok {
			words[i] = newWord
		}
	}
	return strings.Join(words, " ")
}
