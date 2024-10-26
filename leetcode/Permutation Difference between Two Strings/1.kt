class Solution {
    fun findPermutationDifference(s: String, t: String): Int {
        val m: MutableMap<Char, Int> = mutableMapOf()
        for ((i, c) in s.withIndex()) {
            m[c] = i
        }
        var ans = 0
        for ((i, c) in t.withIndex()) {
            val diff = m[c]!! - i
            ans += abs(diff)
        }
        return ans
    }
}