func hammingWeight(num uint32) int {
    var ret int
    for num > 0 {
        ret += int(num % 2)
        num = num >> 1
    }
    return ret;
}
