// https://leetcode.com/problems/palindrome-number/

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	ret := 0
	mod := 10
	stop := 0
	for stop < 2 {
		if x/mod == 0 {
			stop += 1
		}
		if stop == 2 {
			break
		}
		num := (x % mod) / (mod / 10)
		mod *= 10
		ret *= 10
		ret += num
	}

	return ret == x
}