func maxOperations(nums []int, k int) int {
    dict := map[int]int{}
    for _, n := range nums{
        if _, ok := dict[n]; !ok{
            dict[n] = 1
        }else{
            dict[n]++
        }
    }
    
    ret := 0
    for _, n := range nums{
        v1, ok1 := dict[k-n]
        v2, ok2 := dict[n]
        
        if n == k-n && v1 < 2 {
            ok1 = false
        }
        
        if v1 == 0{
            ok1 = false
        }
        if v2 == 0{
            ok2 = false
        }
        
        if ok1 && ok2{
            dict[n]--
            dict[k-n]--
            ret++
        }
    }
    
    return ret
}
