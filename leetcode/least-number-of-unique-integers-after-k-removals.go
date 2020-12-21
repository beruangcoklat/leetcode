func findLeastNumOfUniqueInts(arr []int, k int) int {
	dict := map[int]int{}
	ret := 0
	for _, a := range arr {
		if _, ok := dict[a]; ok {
			dict[a]++
			continue
		}
		dict[a] = 1
		ret++
	}

	sorted := []int{}
	for _, v := range dict {
		sorted = append(sorted, v)
	}
	sort.Ints(sorted)

	for _, s := range sorted {
		if s > k {
			break
		}
		k -= s
		ret--
		if k <= 0 {
			break
		}
	}

	return ret
}
