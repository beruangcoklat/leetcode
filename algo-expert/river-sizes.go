package program

var col int
var row int
var total int

func ff(matrix [][]int, x, y int) {
	if x < 0 || y < 0 || x >= col || y >= row {
		return
	}

	if matrix[y][x] == 1 {
		total += 1
		matrix[y][x] = 2
		ff(matrix, x+1, y)
		ff(matrix, x-1, y)
		ff(matrix, x, y+1)
		ff(matrix, x, y-1)
	}
}

func RiverSizes(matrix [][]int) []int {
	row = len(matrix)
	col = len(matrix[0])
	ret := []int{}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			total = 0
			if matrix[i][j] == 1 {
				ff(matrix, j, i)
			}
			if total > 0 {
				ret = append(ret, total)
			}
		}
	}

	return ret
}
