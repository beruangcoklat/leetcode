func isValidSudoku(board [][]byte) bool {

	for y := 0; y < 9; y++ {
		rowMap := make(map[byte]struct{})
		for x := 0; x < 9; x++ {
			if board[y][x] == '.' {
				continue
			}
			_, ok := rowMap[board[y][x]]
			if ok {
				fmt.Println(y, x)
				return false
			}
			rowMap[board[y][x]] = struct{}{}
		}
	}

	for x := 0; x < 9; x++ {
		colMap := make(map[byte]struct{})
		for y := 0; y < 9; y++ {
			if board[y][x] == '.' {
				continue
			}
			_, ok := colMap[board[y][x]]
			if ok {
				fmt.Println(y, x)
				return false
			}
			colMap[board[y][x]] = struct{}{}
		}
	}

	for y := 0; y < 3; y++ {
		for x := 0; x < 3; x++ {

			gridMap := make(map[byte]struct{})
			for i := y * 3; i < (y*3)+3; i++ {
				for j := x * 3; j < (x*3)+3; j++ {
					if board[j][i] == '.' {
						continue
					}
					_, ok := gridMap[board[j][i]]
					if ok {
						return false
					}
					gridMap[board[j][i]] = struct{}{}
				}
			}

		}
	}

	return true
}
