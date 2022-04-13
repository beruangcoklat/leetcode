func generateMatrix(n int) [][]int {
	var (
		dirX  = []int{1, 0, -1, 0}
		dirY  = []int{0, 1, 0, -1}
		count = 0
		dir   = 0
		px    = 0
		py    = 0
	)

	ret := make([][]int, n)
	visited := make([][]bool, n)
	for i := 0; i < n; i++ {
		ret[i] = make([]int, n)
		visited[i] = make([]bool, n)
	}

	for {
		count++
		ret[py][px] = count
		visited[py][px] = true

		nx := px + dirX[dir]
		ny := py + dirY[dir]
		if nx < 0 || ny < 0 || nx >= n || ny >= n || visited[ny][nx] {
			dir++
			if dir >= 4 {
				dir = 0
			}
			nx = px + dirX[dir]
			ny = py + dirY[dir]
		}
		if nx < 0 || ny < 0 || nx >= n || ny >= n || visited[ny][nx] {
			break
		}
		px = nx
		py = ny
	}

	return ret
}
