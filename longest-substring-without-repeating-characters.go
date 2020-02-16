// https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	l := len(s)
	if l == 0 {
		return 0
	}
	ret := 1
	for i := 0; i < l; i++ {
		dict := make(map[byte]struct{})
		dict[s[i]] = struct{}{}
		curr := 1
		for j := i + 1; j < l; j++ {
			if _, ok := dict[s[j]]; !ok {
				dict[s[j]] = struct{}{}
				curr += 1
				if curr > ret {
					ret = curr
				}
			} else {
				break
			}
		}
	}
	return ret
}