func reverseWords(s string) string {
	words := []string{}
	currWord := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if currWord != "" {
				words = append(words, currWord)
				currWord = ""
			}
			continue
		}

		currWord += string(s[i])
	}
	if currWord != "" {
		words = append(words, currWord)
	}

	ret := ""
	for i := len(words) - 1; i >= 0; i-- {
		ret += words[i]
		if i > 0 {
			ret += " "
		}
	}
	return ret
}