func transpose(matrix [][]int) [][]int {
	r := len(matrix)
	c := len(matrix[0])
	ret := make([][]int, c)

	for i := 0; i < c; i++ {
		ret[i] = make([]int, r)
	}

	for y := 0; y < c; y++ {
		for x := 0; x < r; x++ {
			ret[y][x] = matrix[x][y]
		}
	}
	return ret
}
