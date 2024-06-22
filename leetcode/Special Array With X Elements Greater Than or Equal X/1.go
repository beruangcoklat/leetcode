type data struct {
	num                 int
	totalCountFromRight int
}

func specialArray(nums []int) int {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}

	datas := make([]data, 0)
	for num, count := range m {
		datas = append(datas, data{
			num:                 num,
			totalCountFromRight: count,
		})
	}

	sort.SliceStable(datas, func(i, j int) bool {
		return datas[i].num < datas[j].num
	})

	for i := len(datas) - 2; i >= 0; i-- {
		datas[i].totalCountFromRight += datas[i+1].totalCountFromRight
	}

	for i := 0; i < len(datas); i++ {
		totalCountFromRight := datas[i].totalCountFromRight
		num := datas[i].num

		if totalCountFromRight == num {
			return totalCountFromRight
		}

		if totalCountFromRight < num && (i == 0 || datas[i-1].num < totalCountFromRight) {
			return totalCountFromRight
		}
	}

	return -1
}
