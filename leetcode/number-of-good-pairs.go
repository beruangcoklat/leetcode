func numIdenticalPairs(nums []int) int {
    ret := 0
    for i := 0 ; i<len(nums) ; i++ {
        for j := i + 1 ; j<len(nums) ; j++{
            if nums[i] == nums[j]{
                ret++
            }
        }
    }
    return ret
}
