func merge(nums1 []int, m int, nums2 []int, n int)  {
    lastIdx1 := m - 1
    lastIdx2 := n - 1

    for i := m + n - 1 ; i >= 0 ; i-- {
        var max int
        if lastIdx2 < 0 || (lastIdx1 >= 0 && nums1[lastIdx1] > nums2[lastIdx2]) {
            max = nums1[lastIdx1]
            lastIdx1--
        } else {
            max = nums2[lastIdx2]
            lastIdx2--
        }
        nums1[i] = max
    }
}
