func process(str string) string {
	ret := ""
	for _, char := range str {
		if char != '#' {
			ret += string(char)
		} else if len(ret) > 0 {
			ret = ret[:len(ret)-1]
		}
	}
	return ret
}

func backspaceCompare(s string, t string) bool {
	return process(s) == process(t)
}