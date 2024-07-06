func intersect(nums1 []int, nums2 []int) []int {
	sort.SliceStable(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})

	sort.SliceStable(nums2, func(i, j int) bool {
		return nums2[i] < nums2[j]
	})

	result := []int{}
	l, r := 0, 0
	for l < len(nums1) && r < len(nums2) {
		if nums1[l] == nums2[r] {
			result = append(result, nums1[l])
			l++
			r++
		} else if nums1[l] < nums2[r] {
			l++
		} else {
			r++
		}
	}
	return result
}