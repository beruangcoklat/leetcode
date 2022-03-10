func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearestValidPoint(x int, y int, points [][]int) int {
	ret := -1
	minDist := -1
	for i := 0; i < len(points); i++ {
		xx := points[i][0]
		yy := points[i][1]

		dist := 0
		if x == xx {
			dist = abs(y - yy)
		} else if y == yy {
			dist = abs(x - xx)
		} else {
			continue
		}

		if minDist == -1 || dist < minDist {
			minDist = dist
			ret = i
		}
	}
	return ret
}
