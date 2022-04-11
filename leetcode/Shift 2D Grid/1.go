func shift(grid [][]int, width, height int) [][]int {
	res := make([][]int, height)
	for i := 0; i < height; i++ {
		res[i] = make([]int, width)
	}
	firstCol := make([]int, height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			idx := x - 1
			if idx < 0 {
				idx = width - 1
			}
			if x == 0 {
				firstCol[y] = grid[y][idx]
			}
			res[y][x] = grid[y][idx]
		}
	}
	for y := 0; y < height; y++ {
		idx := y - 1
		if idx < 0 {
			idx = height - 1
		}
		res[y][0] = firstCol[idx]
	}
	return res
}

func shiftGrid(grid [][]int, k int) [][]int {
	width, height := len(grid[0]), len(grid)
	for i := 0; i < k; i++ {
		grid = shift(grid, width, height)
	}
	return grid
}

