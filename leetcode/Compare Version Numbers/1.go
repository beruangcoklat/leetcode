func convert(str string) int {
	ret := ""
	for _, c := range str {
		if ret == "" && c == '0' {
			continue
		}
		ret += string(c)
	}
	if str == "" {
		return 0
	}
	n, _ := strconv.Atoi(str)
	return n
}

func compareVersion(version1 string, version2 string) int {
	str1 := strings.Split(version1, ".")
	str2 := strings.Split(version2, ".")

	arr1 := []int{}
	arr2 := []int{}

	for _, str := range str1 {
		arr1 = append(arr1, convert(str))
	}

	for _, str := range str2 {
		arr2 = append(arr2, convert(str))
	}

	idx := 0
	for idx < len(arr1) && idx < len(arr2) {
		if arr1[idx] < arr2[idx] {
			return -1
		}
		if arr1[idx] > arr2[idx] {
			return 1
		}
		idx++
	}

	for i := idx; i < len(arr1); i++ {
		if arr1[i] > 0 {
			return 1
		}
	}

	for i := idx; i < len(arr2); i++ {
		if arr2[i] > 0 {
			return -1
		}
	}

	return 0
}
