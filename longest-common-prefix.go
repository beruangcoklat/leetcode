// https://leetcode.com/problems/longest-common-prefix/

func longestCommonPrefix(strs []string) string {
	l := len(strs)
	if l == 0 {
		return ""
	}
	ret := 0
	for i := 0; i < len(strs[0]); i++ {
		same := true
		for j := 0; j < l-1; j++ {
			if len(strs[j]) <= i || len(strs[j+1]) <= i || strs[j][i] != strs[j+1][i] {
				same = false
				break
			}
		}
		if !same {
			break
		}
		ret += 1
	}
	return strs[0][:ret]
}