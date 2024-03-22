func merge(nums1 []int, m int, nums2 []int, n int)  {
    result := make([]int, m + n)
    idx1, idx2, idx := 0, 0, 0
    for idx1 < m && idx2 < n {
        if nums1[idx1] < nums2[idx2] {
            result[idx] = nums1[idx1]
            idx1++
        } else {
            result[idx] = nums2[idx2]
            idx2++
        }
        idx++
    }
    
    for idx1 < m {        
        result[idx] = nums1[idx1]
        idx1++
        idx++
    }
    
    for idx2 < n {        
        result[idx] = nums2[idx2]
        idx2++
        idx++
    }
    
    for i := 0 ; i < m + n ; i++ {
        nums1[i] = result[i]
    }
}
