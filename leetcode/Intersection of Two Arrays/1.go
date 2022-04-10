func intersection(nums1 []int, nums2 []int) []int {
	nums1Map := make(map[int]struct{})
	for _, n := range nums1 {
		nums1Map[n] = struct{}{}
	}

	ret := []int{}
	retMap := make(map[int]struct{})
	for _, n := range nums2 {
		if _, ok := nums1Map[n]; !ok {
			continue
		}
		if _, ok := retMap[n]; ok {
			continue
		}
		ret = append(ret, n)
		retMap[n] = struct{}{}
	}

	return ret
}
