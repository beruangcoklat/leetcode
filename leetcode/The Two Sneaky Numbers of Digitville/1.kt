class Solution {
    fun getSneakyNumbers(nums: IntArray): IntArray {
        val countMap: MutableMap<Int, Int> = mutableMapOf()
        var ans: IntArray = intArrayOf(-1, -1)
        nums.forEach { num ->
            countMap[num] = (countMap[num] ?: 0) + 1
            if (countMap[num] ?: 0 == 1) {
                return@forEach
            }
            if (ans[0] == -1) ans[0] = num else ans[1] = num
        }
        return ans
    }
}