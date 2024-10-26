class Solution {
    fun countConsistentStrings(allowed: String, words: Array<String>): Int {
        var ans = 0
        var allowedMap: MutableMap<Char, Boolean> = mutableMapOf()
        allowed.forEach { c ->
            allowedMap[c] = true
        }

        words.forEach { word ->
            for (c in word) {
                val exists = allowedMap[c] ?: false
                if (!exists) {
                    return@forEach
                }
            }
            ans += 1
        }
        return ans
    }
}