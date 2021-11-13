var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{-1, 1, 0, 0}
)

type point struct {
	x, y int
}

func inBorder(x, y int, board [][]byte) bool {
	return x == 0 || y == 0 || x == len(board[0])-1 || y == len(board)-1
}

func bfs(x, y int, board [][]byte, visited [][]bool) []point {
	ret := []point{}
	list := []point{{x, y}}
	for len(list) > 0 {
		curr := list[0]
		list = list[1:]

		if curr.x < 0 || curr.x >= len(board[0]) ||
			curr.y < 0 || curr.y >= len(board) ||
			visited[curr.y][curr.x] ||
			board[curr.y][curr.x] == 'X' {
			continue
		}

		visited[curr.y][curr.x] = true
		ret = append(ret, curr)

		for i := 0; i < len(dirX); i++ {
			nx := curr.x + dirX[i]
			ny := curr.y + dirY[i]
			list = append(list, point{nx, ny})
		}
	}
	return ret
}

func canUpdate(points []point, board [][]byte) bool {
	mark := make(map[string]struct{})
	for _, point := range points {
		key := fmt.Sprintf("%d-%d", point.x, point.y)
		mark[key] = struct{}{}

		if board[point.y][point.x] == 'O' && inBorder(point.x, point.y, board) {
			return false
		}
	}

	for _, point := range points {
		for i := 0; i < len(dirX); i++ {
			nx := point.x + dirX[i]
			ny := point.y + dirY[i]

			if nx < 0 || ny < 0 || nx >= len(board[0]) || ny >= len(board) {
				continue
			}

			key := fmt.Sprintf("%d-%d", nx, ny)
			if _, ok := mark[key]; ok {
				continue
			}

			if board[ny][nx] == 'O' {
				return false
			}
		}
	}

	return true
}

func solve(board [][]byte) {
	visited := make([][]bool, len(board))
	for i := 0; i < len(board); i++ {
		visited[i] = make([]bool, len(board[i]))
	}

	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] == 'X' {
				continue
			}
			points := bfs(x, y, board, visited)
			if !canUpdate(points, board) {
				continue
			}

			for _, point := range points {
				board[point.y][point.x] = 'X'
			}
		}
	}
}
