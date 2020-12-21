func runningSum(nums []int) []int {
    ret := []int{}
    curr := 0
    for i:= 0 ; i<len(nums) ; i++{
        curr += nums[i]
        ret = append(ret, curr)
    }
    return ret
}
