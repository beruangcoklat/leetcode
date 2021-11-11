var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{-1, 1, 0, 0}
)

func dfs(x, y int, grid [][]byte) {
	if x < 0 || y < 0 || y >= len(grid) || x >= len(grid[0]) {
		return
	}
	if grid[y][x] == '0' || grid[y][x] == '2' {
		return
	}
	grid[y][x] = '2'
	for i := 0; i < 4; i++ {
		dfs(x+dirX[i], y+dirY[i], grid)
	}
}

func numIslands(grid [][]byte) int {
	ret := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			if grid[y][x] == '1' {
				ret++
				dfs(x, y, grid)
			}
		}
	}
	return ret
}
