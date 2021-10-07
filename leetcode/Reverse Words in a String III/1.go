func reverseString(s []byte) {
	l, r := 0, len(s)-1
	for l < r {
		temp := s[l]
		s[l] = s[r]
		s[r] = temp
		l++
		r--
	}
}

func reverseWords(s string) string {
	splits := strings.Split(s, " ")
	reversedSplits := make([]string, len(splits))

	for i, str := range splits {
		bytes := []byte(str)
		reverseString(bytes)
		reversedSplits[i] = string(bytes)
	}

	return strings.Join(reversedSplits, " ")
}
