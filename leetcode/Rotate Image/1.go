func transpose(matrix [][]int) [][]int {
	r, c := len(matrix), len(matrix[0])
	result := make([][]int, r)
	for i := 0; i < r; i++ {
		result[i] = make([]int, c)
	}
	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			result[y][x] = matrix[x][y]
		}
	}
	return result
}

func reverse(matrix [][]int) {
	r, c := len(matrix), len(matrix[0])
	for y := 0; y < r; y++ {
		for x := 0; x < c/2; x++ {
			matrix[y][x], matrix[y][c-1-x] = matrix[y][c-1-x], matrix[y][x]
		}
	}
}

func rotate(matrix [][]int) {
	result := transpose(matrix)
	reverse(result)
	r, c := len(matrix), len(matrix[0])
	for y := 0; y < r; y++ {
		for x := 0; x < c; x++ {
			matrix[y][x] = result[y][x]
		}
	}
}
