func createArr(str string) []int {
	arr := make([]int, 26)
	for _, s := range str {
		arr[s-'a']++
	}
	return arr
}

func isSame(arr1, arr2 []int) bool {
	for i := 0; i < 26; i++ {
		if arr1[i] != arr2[i] {
			return false
		}
	}
	return true
}

func checkInclusion(s1 string, s2 string) bool {
	s1Arr := createArr(s1)

	for i := 0; i <= len(s2)-len(s1); i++ {
		str := s2[i : i+len(s1)]
		if isSame(s1Arr, createArr(str)) {
			return true
		}
	}

	return false
}