func getMin(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func getMax(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func maxArea(height []int) int {
	lenHeight := len(height)
	max := -1
    left, right := 0, lenHeight - 1
    for left < right {
        curr := (right - left)  * getMin(height[left], height[right])
        max = getMax(max, curr)
        if height[left] < height[right] {
            left++
        } else {
            right--
        }
    }
	return max
}
