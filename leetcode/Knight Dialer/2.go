var (
	phone = [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
		{-1, 0, -1},
	}
	dirX = []int{1, -1, 1, -1, 2, 2, -2, -2}
	dirY = []int{2, 2, -2, -2, 1, -1, 1, -1}
	mod  = 1000000007

	cache [5001][5][5]int
)

func solve(n int, x, y int) int {
	if n == 0 {
		return 1
	}

	if cache[n][x][y] != 0 {
		return cache[n][x][y]
	}

	result := 0
	for i := 0; i < len(dirX); i++ {
		nextX := x + dirX[i]
		nextY := y + dirY[i]
		if nextX < 0 || nextX >= len(phone[0]) || nextY < 0 ||
			nextY >= len(phone) || phone[nextY][nextX] == -1 {
			continue
		}
		result += solve(n-1, nextX, nextY)
		result %= mod
	}

	result %= mod

	cache[n][x][y] = result
	return result
}

func knightDialer(n int) int {
	cache = [5001][5][5]int{}
	result := 0
	for y := 0; y < len(phone); y++ {
		for x := 0; x < len(phone[y]); x++ {
			if phone[y][x] == -1 {
				continue
			}
			result += solve(n-1, x, y)
			result %= mod
		}
	}
	return result % mod
}