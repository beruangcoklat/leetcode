class Solution {
    fun findWordsContaining(words: Array<String>, x: Char): List<Int> {
        var ans: MutableList<Int> = mutableListOf()
        words.forEachIndexed { index, word ->
            if (word.contains(x)) {
                ans.add(index)
            }
        }
        return ans
    }
}