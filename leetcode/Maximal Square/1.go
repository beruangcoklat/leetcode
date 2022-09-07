func min(arr ...int) int {
	ret := arr[0]
	for i := 1; i < len(arr); i++ {
		if ret > arr[i] {
			ret = arr[i]
		}
	}
	return ret
}

func maximalSquare(matrix [][]byte) int {
	r := len(matrix)
	c := len(matrix[0])
	dp := make([][]int, r)

	for y := 0; y < r; y++ {
		dp[y] = make([]int, c)
		for x := 0; x < c; x++ {
			dp[y][x] = int(matrix[y][x] - '0')
		}
	}

	max := -1
	haveOne := false
	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			if matrix[y][x] == '0' {
				continue
			}
			haveOne = true

			if y == 0 || x == 0 {
				continue
			}

			dp[y][x] = min(dp[y][x-1], dp[y-1][x], dp[y-1][x-1]) + 1
			if dp[y][x] > max {
				max = dp[y][x]
			}
		}
	}

	if !haveOne {
		return 0
	}

	if haveOne && max == 0 {
		max = 1
	}

	return max * max
}
