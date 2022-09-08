func validCoordinate(x, y, r, c int) bool {
	return x >= 0 && x < c && y >= 0 && y < r
}

func diagonalSort(mat [][]int) [][]int {
	r := len(mat)
	c := len(mat[0])

	datas := [][]int{}

	for x := c - 1; x >= 0; x-- {
		coorY := 0
		coorX := x
		data := []int{}
		for validCoordinate(coorX, coorY, r, c) {
			data = append(data, mat[coorY][coorX])
			coorX++
			coorY++
		}
		datas = append(datas, data)
	}

	for y := 0; y < r; y++ {
		coorY := y
		coorX := 0
		data := []int{}
		for validCoordinate(coorX, coorY, r, c) {
			data = append(data, mat[coorY][coorX])
			coorX++
			coorY++
		}
		datas = append(datas, data)
	}

	for i := 0; i < len(datas); i++ {
		sort.SliceStable(datas[i], func(a, b int) bool {
			return datas[i][a] < datas[i][b]
		})
	}

	idxData := 0
	for x := c - 1; x >= 0; x-- {
		data := datas[idxData]
		idxData++

		coorY := 0
		coorX := x

		for _, a := range data {
			mat[coorY][coorX] = a
			coorY++
			coorX++
		}
	}

	for y := 0; y < r; y++ {
		data := datas[idxData]
		idxData++

		coorY := y
		coorX := 0

		for _, a := range data {
			mat[coorY][coorX] = a
			coorY++
			coorX++
		}
	}

	return mat
}
