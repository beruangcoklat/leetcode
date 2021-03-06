var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
		return []string{}
	}

	chars := phoneMap[digits[:1]]

	if len(digits) == 1 {
		ret := []string{}
		for i := 0; i < len(chars); i++ {
			ret = append(ret, chars[i:i+1])
		}
		return ret
	}

	ret := []string{}

	for i := 0; i < len(chars); i++ {
		res := letterCombinations(digits[1:])
		for j := 0; j < len(res); j++ {
			curr := chars[i:i+1] + res[j]
			ret = append(ret, curr)
		}
	}

	return ret
}