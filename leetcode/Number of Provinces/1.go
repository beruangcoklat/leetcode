func findCircleNum(isConnected [][]int) int {
	citiesCount := len(isConnected)
	visited := make(map[int]struct{})
	ret := 0

	for aaa := 0; aaa < citiesCount; aaa++ {
		if _, ok := visited[aaa]; ok {
			continue
		}

		ret++
		list := []int{aaa}
		for len(list) > 0 {
			curr := list[0]
			list = list[1:]

			if _, ok := visited[curr]; ok {
				continue
			}
			visited[curr] = struct{}{}

			for i := 0; i < citiesCount; i++ {
				if curr == i || isConnected[curr][i] == 0 {
					continue
				}
				list = append(list, i)
			}
		}
	}
	return ret
}
