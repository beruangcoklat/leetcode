func numberOfBeams(bank []string) int {
	devicePerFloorCount := []int{}
	for _, row := range bank {
		count := 0
		for _, c := range row {
			if c == '1' {
				count++
			}
		}
		if count > 0 {
			devicePerFloorCount = append(devicePerFloorCount, count)
		}
	}

	ret := 0
	for i := 0; i < len(devicePerFloorCount)-1; i++ {
		ret += devicePerFloorCount[i] * devicePerFloorCount[i+1]
	}
	return ret
}
