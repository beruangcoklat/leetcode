var (
	dirX = []int{0, 0, 1, -1}
	dirY = []int{1, -1, 0, 0}
)

func _floodFill(image [][]int, sr int, sc int, newColor int, flagMap map[string]struct{}) {
	oldColor := image[sr][sc]
	image[sr][sc] = newColor

	key := fmt.Sprintf("%v-%v", sr, sc)
	if _, ok := flagMap[key]; ok {
		return
	}
	flagMap[key] = struct{}{}

	for i := 0; i < len(dirX); i++ {
		nx := sc + dirX[i]
		ny := sr + dirY[i]
		if nx < 0 || ny < 0 || nx >= len(image[0]) || ny >= len(image) {
			continue
		}
		if image[ny][nx] != oldColor {
			continue
		}
		_floodFill(image, ny, nx, newColor, flagMap)
	}
}

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	_floodFill(image, sr, sc, newColor, make(map[string]struct{}))
	return image
}
