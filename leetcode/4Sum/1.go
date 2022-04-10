func check(numMap map[int]int, nums []int) bool {
	m := make(map[int]int)
	for _, n := range nums {
		m[n]++
	}
	for k, v := range m {
		if numMap[k] < v {
			return false
		}
	}
	return true
}

func makeKey(nums []int) string {
	ret := ""
	done := false
	for i := 0; i < 3; i++ {
		if !done && nums[3] < nums[i] {
			done = true
			ret += strconv.Itoa(nums[3]) + "_"
		}
		ret += strconv.Itoa(nums[i]) + "_"
	}
	if !done {
		ret += strconv.Itoa(nums[3]) + "_"
	}
	return ret
}

func fourSum(nums []int, target int) [][]int {
	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	ret := [][]int{}
	retMap := make(map[string]struct{})

	for i := 0; i < len(nums); i++ {
		a := nums[i]
		for ii := i + 1; ii < len(nums); ii++ {
			b := nums[ii]
			for iii := ii + 1; iii < len(nums); iii++ {
				c := nums[iii]
				d := target - a - b - c
				arr := []int{a, b, c, d}
				if !check(numMap, arr) {
					continue
				}
				key := makeKey(arr)
				if _, ok := retMap[key]; ok {
					continue
				}
				ret = append(ret, arr)
				retMap[key] = struct{}{}
			}
		}
	}

	return ret
}
