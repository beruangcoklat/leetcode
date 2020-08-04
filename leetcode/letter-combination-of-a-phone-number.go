// https://leetcode.com/problems/letter-combinations-of-a-phone-number

func solve(s string, dict map[string]string, curr string) []string {
	if s == "" || len(s) == 0 {
		return []string{curr}
	}

	ret := []string{}
	for _, char := range dict[s[0:1]] {
		ret = append(ret, solve(s[1:], dict, curr+string(char))...)
	}
	return ret
}

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}
	dict := make(map[string]string)
	dict["2"] = "abc"
	dict["3"] = "def"
	dict["4"] = "ghi"
	dict["5"] = "jkl"
	dict["6"] = "mno"
	dict["7"] = "pqrs"
	dict["8"] = "tuv"
	dict["9"] = "wxyz"

	return solve(digits, dict, "")
}

