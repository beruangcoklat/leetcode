var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{1, -1, 0, 0}
)

func search(x, y int, board [][]byte, word string) bool {
	if word[0] != board[y][x] {
		return false
	}
	if len(word) == 1 {
		return true
	}

	temp := board[y][x]
	board[y][x] = ' '

	nextWord := word[1:]
	for i := 0; i < 4; i++ {
		nx := x + dirX[i]
		ny := y + dirY[i]
		if nx < 0 || ny < 0 || nx >= len(board[0]) || ny >= len(board) {
			continue
		}
		if search(nx, ny, board, nextWord) {
			return true
		}
	}

	board[y][x] = temp
	return false
}

func exist(board [][]byte, word string) bool {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if board[y][x] != word[0] {
				continue
			}
			if search(x, y, board, word) {
				return true
			}
		}
	}
	return false
}
