var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{-1, 1, 0, 0}
)

type point struct {
	x, y int
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

			canUpdate := true
			for _, point := range points {
				if point.x == 0 || point.y == 0 || point.x == len(board[0])-1 || point.y == len(board)-1 {
					canUpdate = false
					break
				}
			}
			if !canUpdate {
				continue
			}

			for _, point := range points {
				board[point.y][point.x] = 'X'
			}
		}
	}
}
