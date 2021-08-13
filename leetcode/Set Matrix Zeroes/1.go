func update(x, y int, matrix [][]int) {
	for i := 0; i < len(matrix[y]); i++ {
		matrix[y][i] = 0
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i][x] = 0
	}
}

func setZeroes(matrix [][]int) {
	originalMatrix := [][]int{}
	for y := 0; y < len(matrix); y++ {
		rows := []int{}
		for x := 0; x < len(matrix[y]); x++ {
			rows = append(rows, matrix[y][x])
		}
		originalMatrix = append(originalMatrix, rows)
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if originalMatrix[y][x] == 0 {
				update(x, y, matrix)
			}
		}
	}
}
