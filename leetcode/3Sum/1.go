func swap(arr []int, a, b int) {
	temp := arr[a]
	arr[a] = arr[b]
	arr[b] = temp
}

func makeKey(arr []int) string {
	if arr[0] > arr[1] {
		swap(arr, 0, 1)
	}
	if arr[1] > arr[2] {
		swap(arr, 1, 2)
	}
	if arr[0] > arr[1] {
		swap(arr, 0, 1)
	}

	key := ""
	for _, a := range arr {
		key += fmt.Sprintf("%v_", a)
	}
	return key
}

func threeSum(nums []int) [][]int {
	numMap := make(map[int]int)
	for _, num := range nums {
		numMap[num]++
	}

	resultMap := make(map[string]struct{})

	ret := [][]int{}
	for i := 0; i < len(nums); i++ {
		a := nums[i]
		numMap[a]--

		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			c := -(a + b)

			if b == c && numMap[b] < 2 {
				continue
			}

			if numMap[b] < 1 || numMap[c] < 1 {
				continue
			}

			arr := []int{a, b, c}
			key := makeKey(arr)

			if _, ok := resultMap[key]; ok {
				continue
			}
			resultMap[key] = struct{}{}
			ret = append(ret, arr)
		}
	}

	return ret
}
