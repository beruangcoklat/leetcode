var (
	dirX = []int{0, 0, 1, -1, 1, 1, -1, -1}
	dirY = []int{1, -1, 0, 0, 1, -1, 1, -1}
)

func getAliveNeighbors(board [][]int, x, y int) int {
	ret := 0
	for i := 0; i < len(dirX); i++ {
		nx := x + dirX[i]
		ny := y + dirY[i]
		if nx < 0 || ny < 0 || nx >= len(board[0]) || ny >= len(board) {
			continue
		}
		ret += board[ny][nx]
	}
	return ret
}

func gameOfLife(board [][]int) {
	cloneBoard := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		cloneBoard[i] = make([]int, len(board[0]))
		for j := 0; j < len(board[0]); j++ {
			cloneBoard[i][j] = board[i][j]
		}
	}
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			an := getAliveNeighbors(cloneBoard, x, y)

			if board[y][x] == 1 && an < 2 {
				board[y][x] = 0
				continue
			}

			if board[y][x] == 1 && an >= 2 && an <= 3 {
				continue
			}

			if board[y][x] == 1 && an >= 2 && an > 3 {
				board[y][x] = 0
				continue
			}

			if board[y][x] == 0 && an == 3 {
				board[y][x] = 1
				continue
			}

		}
	}
}
