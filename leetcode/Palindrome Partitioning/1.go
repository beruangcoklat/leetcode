var result [][]string

func partition(s string) [][]string {
	result = make([][]string, 0)
	solve(s, 0, []string{}, 0)
	return result
}

func solve(s string, idx int, localResult []string, totalLength int) {
	if totalLength == len(s) {
		result = append(result, localResult)
		return
	}

	for length := 1; idx+length <= len(s); length++ {
		str := s[idx : idx+length]
		if isPalindrome(str) {
			solve(s, idx+length, clone(localResult, str), totalLength+len(str))
		}
	}
}

func isPalindrome(str string) bool {
	length := len(str)
	for i := 0; i < length/2; i++ {
		if str[i] != str[length-1-i] {
			return false
		}
	}
	return true
}

func clone(str []string, s string) []string {
	res := make([]string, len(str)+1)
	copy(res, str)
	res[len(str)] = s
	return res
}
