// https://leetcode.com/problems/median-of-two-sorted-arrays/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	idx1 := 0
	idx2 := 0
	l1 := len(nums1)
	l2 := len(nums2)
	lll := l1 + l2
	list := []int{}

	for idx1 < l1 || idx2 < l2 {
		a1 := -100
		a2 := -100

		if idx1 < l1 {
			a1 = nums1[idx1]
			if idx2 >= l2 {
				list = append(list, a1)
				idx1 += 1
				continue
			}
		}

		if idx2 < l2 {
			a2 = nums2[idx2]
			if idx1 >= l1 {
				list = append(list, a2)
				idx2 += 1
				continue
			}
		}

		if a1 < a2 {
			list = append(list, a1)
			idx1 += 1
		} else {
			list = append(list, a2)
			idx2 += 1
		}
	}
	if lll%2 == 1 {
		return float64(list[lll/2])
	}

	x := float64(list[lll/2])
	y := float64(list[lll/2-1])
	return float64((x + y) / 2)
}