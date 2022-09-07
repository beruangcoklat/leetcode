func transpose(matrix [][]int) {
	r, c := len(matrix), len(matrix[0])
	for y := 0; y < r; y++ {
		for x := y + 1; x < c; x++ {
			matrix[y][x], matrix[x][y] = matrix[x][y], matrix[y][x]
		}
	}
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
	transpose(matrix)
	reverse(matrix)
}
