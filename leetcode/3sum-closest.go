// https://leetcode.com/problems/3sum-closest

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func threeSumClosest(nums []int, target int) int {
	lenN := len(nums)
	min := 0
	ret := 0
	first := true
	for a := 0; a < lenN; a++ {
		for b := a + 1; b < lenN; b++ {
			for c := b + 1; c < lenN; c++ {
				d := nums[a] + nums[b] + nums[c]
				e := abs(target - d)
				if e < min || first {
					min = e
					ret = d
					first = false
				}
			}
		}
	}
	return ret
}
