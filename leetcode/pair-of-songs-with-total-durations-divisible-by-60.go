func numPairsDivisibleBy60(time []int) int {
    ret := 0
    for i := 0 ; i<len(time) ; i++{
        for j := i+1 ; j<len(time) ; j++{
            if (time[i] + time[j]) % 60 == 0{
                ret++
            }
        }
    }
    return ret
}
