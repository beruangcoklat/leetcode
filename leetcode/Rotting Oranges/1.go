type point struct {
	x int
	y int
}

func orangesRotting(grid [][]int) int {
	height := len(grid)
	width := len(grid[0])

	points := []point{}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 2 {
				points = append(points, point{x, y})
			}
		}
	}

	dirX := []int{0, 0, 1, -1}
	dirY := []int{-1, 1, 0, 0}
	move := 0

	for len(points) > 0 {
		size := len(points)
		temp := false
		for i := 0; i < size; i++ {
			pts := points[0]
			points = points[1:]

			for i := 0; i < len(dirX); i++ {
				nx := pts.x + dirX[i]
				ny := pts.y + dirY[i]

				if nx < 0 || ny < 0 || nx >= width || ny >= height {
					continue
				}

				if grid[ny][nx] == 0 || grid[ny][nx] == 2 {
					continue
				}

				grid[ny][nx] = 2
				points = append(points, point{nx, ny})
				temp = true
			}
		}

		if !temp {
			break
		}
		move++
	}

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if grid[y][x] == 1 {
				return -1
			}
		}
	}

	return move
}
