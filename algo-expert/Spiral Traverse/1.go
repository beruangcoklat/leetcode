package main

import "fmt"

func SpiralTraverse(array [][]int) []int {
	dirX := []int{1, 0, -1, 0}
	dirY := []int{0, 1, 0, -1}
	dir := 0
	currX, currY := 0, 0
	visitedMap := map[string]struct{}{}

	ret := []int{}

	r := len(array)
	c := len(array[0])
	count := r * c

	for count > 0 {
		count--
		ret = append(ret, array[currY][currX])
		visitedMap[fmt.Sprintf("%d-%d", currX, currY)] = struct{}{}

		nextX := currX + dirX[dir]
		nextY := currY + dirY[dir]

		_, visited := visitedMap[fmt.Sprintf("%d-%d", nextX, nextY)]
		if nextX < 0 || nextY < 0 || nextX >= c || nextY >= r || visited {
			dir = (dir + 1) % 4
		}

		nextX = currX + dirX[dir]
		nextY = currY + dirY[dir]
		currX = nextX
		currY = nextY
	}

	return ret
}
