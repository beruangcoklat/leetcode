var cache = map[string]int{}

func minPathSum(grid [][]int) int {
	cache = make(map[string]int)
	return solve(0, 0, grid)
}

func solve(x, y int, grid [][]int) int {
	maxY := len(grid)
	maxX := len(grid[0])

	if x == maxX-1 && y == maxY-1 {
		return grid[y][x]
	}

	key := fmt.Sprintf("%d-%d", x, y)
	if v, ok := cache[key]; ok {
		return v
	}

	first := true
	min := 0
	if x+1 < maxX {
		a := solve(x+1, y, grid) + grid[y][x]
		if first || a < min {
			first = false
			min = a
		}
	}
	if y+1 < maxY {
		a := solve(x, y+1, grid) + grid[y][x]
		if first || a < min {
			first = false
			min = a
		}
	}

	cache[key] = min
	return min
}
