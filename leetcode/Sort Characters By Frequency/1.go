type data struct {
	val   rune
	count int
}

func frequencySort(s string) string {
	countMap := make(map[rune]int)
	for _, c := range s {
		countMap[c]++
	}

	datas := []data{}
	for k, v := range countMap {
		datas = append(datas, data{
			val:   k,
			count: v,
		})
	}

	sort.SliceStable(datas, func(i, j int) bool {
		return datas[i].count > datas[j].count
	})

	ret := ""
	for _, data := range datas {
		for i := 0; i < data.count; i++ {
			ret += string(data.val)
		}
	}
	return ret
}
