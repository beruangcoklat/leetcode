class Solution {
    fun leftRightDifference(nums: IntArray): IntArray {
        val left = IntArray(nums.size)
        val right = IntArray(nums.size)
        val ans = IntArray(nums.size)

        val sum = nums.sum()
        var leftSum = sum
        var rightSum = sum

        for (i in nums.size - 1 downTo 0) {
            leftSum -= nums[i]
            left[i] = leftSum
        }

        for (i in nums.indices) {
            rightSum -= nums[i]
            right[i] = rightSum
        }

        for (i in nums.indices) {
            ans[i] = abs(left[i] - right[i])
        }

        return ans
    }
}