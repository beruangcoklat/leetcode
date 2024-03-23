func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	str := ""
	for _, c := range s {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			str += string(c)
		}
	}

	size := len(str)
	for i := 0; i < size/2; i++ {
		if str[i] != str[size-1-i] {
			return false
		}
	}
	return true
}
