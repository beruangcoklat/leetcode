func reversePrefix(word string, ch byte) string {
	idx := -1
	for i := 0; i < len(word); i++ {
		if ch == word[i] {
			idx = i
			break
		}
	}
	if idx == -1 {
		return word
	}
	return reverse(word[:idx+1]) + word[idx+1:]
}

func reverse(str string) string {
	ret := ""
	for i := len(str) - 1; i >= 0; i-- {
		ret += string(str[i])
	}
	return ret
}
