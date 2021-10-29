var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{-1, 1, 0, 0}
	curr = 0
)

func getArea(x, y int, grid [][]int, flag [][]bool) {
	if flag[y][x] {
		return
	}

	curr++
	flag[y][x] = true

	for i := 0; i < len(dirX); i++ {
		nx := x + dirX[i]
		ny := y + dirY[i]

		if nx < 0 || ny < 0 || nx >= len(grid[0]) || ny >= len(grid) {
			continue
		}

		if grid[ny][nx] == 0 {
			continue
		}

		getArea(nx, ny, grid, flag)
	}
}

func maxAreaOfIsland(grid [][]int) int {
	width := len(grid[0])
	height := len(grid)

	flag := make([][]bool, height)
	for i := 0; i < height; i++ {
		flag[i] = make([]bool, width)
	}

	max := 0

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 0 || flag[y][x] {
				continue
			}

			curr = 0
			getArea(x, y, grid, flag)
			if curr > max {
				max = curr
			}
		}
	}

	return max
}
