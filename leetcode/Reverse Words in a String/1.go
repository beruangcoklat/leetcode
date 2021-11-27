func reverseWords(s string) string {
	splits := strings.Fields(s)
	reversed := make([]string, len(splits))
	for i := 0; i < len(splits); i++ {
		reversed[i] = splits[len(splits)-1-i]
	}
	return strings.Join(reversed, " ")
}
