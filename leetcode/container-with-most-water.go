// https://leetcode.com/problems/container-with-most-water/

func maxArea(height []int) int {
	lenHeight := len(height)
	max := -1
	for i := 0; i < lenHeight; i++ {
		for j := i + 1; j < lenHeight; j++ {
			dist := j - i
			minHeight := height[i]
			if height[j] < minHeight {
				minHeight = height[j]
			}
			curr := dist * minHeight
			if curr > max || max == -1 {
				max = curr
			}
		}
	}
	return max
}
