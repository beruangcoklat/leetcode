func strStr(haystack string, needle string) int {
	needleLength := len(needle)
	for i := 0; i <= len(haystack)-needleLength; i++ {
		if haystack[i:i+needleLength] == needle {
			return i
		}
	}
	return -1
}
