var result = []string{}

func recur(idx int, s string) {
	if idx == len(s) {
		result = append(result, s)
		return
	}

	if !unicode.IsLetter(rune(s[idx])) {
		recur(idx+1, s)
		return
	}

	recur(idx+1, fmt.Sprintf("%v%v%v", s[:idx], string(unicode.ToUpper(rune(s[idx]))), s[idx+1:]))
	recur(idx+1, fmt.Sprintf("%v%v%v", s[:idx], string(unicode.ToLower(rune(s[idx]))), s[idx+1:]))
}

func letterCasePermutation(s string) []string {
	result = []string{}
	recur(0, s)
	return result
}
