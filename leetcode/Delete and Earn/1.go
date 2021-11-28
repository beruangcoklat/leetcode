var memo map[int]int

func solve(idx int, arr []int, freqMap map[int]int) int {
	if idx >= len(arr) {
		return 0
	}

	if val, ok := memo[idx]; ok {
		return val
	}

	notTake := solve(idx+1, arr, freqMap)
	max := notTake

	take := solve(func() int {
		if idx+1 < len(arr) && arr[idx+1] == arr[idx]+1 {
			return idx + 2
		}
		return idx + 1
	}(), arr, freqMap) + (freqMap[arr[idx]] * arr[idx])
	if take > max {
		max = take
	}

	memo[idx] = max
	return max
}

func deleteAndEarn(nums []int) int {
	memo = make(map[int]int)
	freqMap := make(map[int]int)
	for _, n := range nums {
		freqMap[n]++
	}
	arr := []int{}
	for k := range freqMap {
		arr = append(arr, k)
	}
	sort.SliceStable(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	return solve(0, arr, freqMap)
}
