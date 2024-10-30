
class Solution {
    fun movesToMakeZigzag(nums: IntArray): Int {
        val a = solve(nums.clone(), true)
        val b = solve(nums.clone(), false)
        return min(a, b)
    }

    private fun solve(nums: IntArray, isEven: Boolean): Int {
        var ans = 0
        val start = if (isEven) 0 else 1
        for (i in start..<nums.size step 2) {
            if (i - 1 >= 0 && nums[i] <= nums[i - 1]) {
                ans += getDiff(nums[i - 1], nums[i] - 1)
                nums[i - 1] = nums[i] - 1
            }
            if (i + 1 < nums.size && nums[i] <= nums[i + 1]) {
                ans += getDiff(nums[i + 1], nums[i] - 1)
                nums[i + 1] = nums[i] - 1
            }
        }
        return ans
    }

    private fun getDiff(from: Int, to: Int): Int {
        return abs(to - from)
    }
}