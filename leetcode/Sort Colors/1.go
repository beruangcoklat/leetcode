func sortColors(nums []int) {
	var a, b, c []int
	for _, n := range nums {
		switch n {
		case 0:
			a = append(a, n)
		case 1:
			b = append(b, n)
		case 2:
			c = append(c, n)
		}
	}
	result := append(a, b...)
	result = append(result, c...)
	copy(nums, result)
}