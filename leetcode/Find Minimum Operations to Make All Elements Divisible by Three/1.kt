class Solution {
    fun minimumOperations(nums: IntArray): Int {
        var ans = 0
        for (num in nums) {
            var mod = num % 3
            ans += if (mod != 0) 1 else 0
        }
        return ans
    }
}