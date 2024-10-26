class Solution {
    fun numberOfEmployeesWhoMetTarget(hours: IntArray, target: Int): Int {
        var ans = 0
        hours.forEach { hour ->
            if (hour >= target) {
                ans++
            }
        }
        return ans
    }
}