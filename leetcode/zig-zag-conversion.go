// https://leetcode.com/problems/zigzag-conversion/

func getChar(s string, lenS int, idx int, dict map[int]struct{}) (string, bool) {
	if idx >= lenS || idx < 0 {
		return "", false
	}
	if _, ok := dict[idx]; ok {
		return "", false
	}
	return s[idx : idx+1], true
}

func convert(s string, numRows int) string {
	dict := make(map[int]struct{})
	increment := (numRows * 2) - 2
	lenS := len(s)
	ret := ""
	for i := 0; i < lenS; i++ {
		idx := i - increment
		for true {
			idx += increment
			char, valid := getChar(s, lenS, idx, dict)
			if !valid {
				break
			}
			ret += char
			dict[idx] = struct{}{}

			if i > 0 && i < lenS-1 {
				idx2 := idx + increment - (i * 2)
				char, valid = getChar(s, lenS, idx2, dict)
				if !valid {
					break
				}
				ret += char
				dict[idx2] = struct{}{}
			}
		}
	}
	return ret
}
