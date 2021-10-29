type point struct {
	x int
	y int
}

var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{-1, 1, 0, 0}
)

func updateMatrix(mat [][]int) [][]int {
	width := len(mat[0])
	height := len(mat)

	distMat := make([][]int, height)
	for i := 0; i < height; i++ {
		distMat[i] = make([]int, width)
	}

	points := []point{}
	nextPoints := []point{}
	nextPointsMap := make(map[string]struct{})

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if mat[y][x] == 0 {
				points = append(points, point{x, y})
			}
		}
	}

	for len(points) > 0 {
		for _, pts := range points {
			x := pts.x
			y := pts.y
			for i := 0; i < len(dirX); i++ {
				nx := x + dirX[i]
				ny := y + dirY[i]
				if nx < 0 || ny < 0 || nx >= width || ny >= height {
					continue
				}
				if mat[ny][nx] == 0 {
					continue
				}
				if distMat[ny][nx] != 0 {
					continue
				}
				key := fmt.Sprintf("%v-%v", nx, ny)
				if _, ok := nextPointsMap[key]; ok {
					continue
				}

				distMat[ny][nx] = distMat[y][x] + 1
				nextPoints = append(nextPoints, point{nx, ny})
				nextPointsMap[key] = struct{}{}
			}
		}

		points = nextPoints
		nextPoints = []point{}
		nextPointsMap = make(map[string]struct{})
	}

	return distMat
}
