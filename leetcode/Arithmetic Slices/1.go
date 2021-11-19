func numberOfArithmeticSlices(nums []int) int {
    ret := 0
    diff := 0
    curr := 0
    first := true
    
    for i:=2 ; i<len(nums) ; i++{
        if nums[i] - nums[i-1] == nums[i-1] - nums[i-2] {
            currDiff := nums[i] - nums[i-1]
            if first || diff != currDiff {
                curr = 1
                first = false
            } else {
                curr++
            }
            ret += curr
            diff = currDiff
        } else {
            first = true
            curr = 1
        }
    }
    return ret
}
