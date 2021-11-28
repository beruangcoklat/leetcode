func spiralOrder(matrix [][]int) []int {
	dirX := []int{1, 0, -1, 0}
	dirY := []int{0, 1, 0, -1}
	dir := 0

	width := len(matrix[0])
	height := len(matrix)

	total := width * height
	ret := []int{}
	x, y := 0, 0

	visited := make([][]bool, height)
	for i := range visited {
		visited[i] = make([]bool, width)
	}

	for len(ret) < total {
		visited[y][x] = true
		ret = append(ret, matrix[y][x])

		nx := x + dirX[dir]
		ny := y + dirY[dir]
		if nx >= width || nx < 0 || ny >= height || ny < 0 || visited[ny][nx] {
			dir++
			if dir >= len(dirX) {
				dir = 0
			}
		}
		x += dirX[dir]
		y += dirY[dir]
	}

	return ret
}
