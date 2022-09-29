type data struct {
	value int
	diff  int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func findClosestElements(arr []int, k int, x int) []int {
	datas := make([]data, len(arr))
	for i, a := range arr {
		datas[i] = data{
			value: a,
			diff:  abs(a - x),
		}
	}
	sort.SliceStable(datas, func(i, j int) bool {
		return datas[i].diff < datas[j].diff
	})

	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[i] = datas[i].value
	}
	sort.SliceStable(ret, func(i, j int) bool {
		return ret[i] < ret[j]
	})
	return ret
}
